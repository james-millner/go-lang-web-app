package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/james-millner/go-lang-web-app/pkg/aws"
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/james-millner/go-lang-web-app/pkg/web"

	"github.com/grokify/html-strip-tags-go"

	"github.com/google/uuid"
)

//ProcessCaseStudyLink function.
func (cs *CaseStudyService) ProcessCaseStudyLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		companyNumber := r.FormValue("company_number")
		url := r.FormValue("url")
		log.Println(companyNumber + " - " + url)

		if url == "" || !web.IsPDFDocument(url) {
			log.Println("Invalid URL received: " + url)
			enc.Encode(&model.CaseStudyDTO{})
			return
		}

		tokens := strings.Split(url, "/")
		fileName := tokens[len(tokens)-1]

		out, oserr := os.Create(fileName)

		if oserr != nil {
			e := fmt.Errorf("Error with creating OS file: %v", oserr)
			log.Fatal(e)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			e := fmt.Errorf("Error with GET request: %v", err)
			log.Fatal(e)
			return
		}

		defer resp.Body.Close()

		io.Copy(out, resp.Body)

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
			return
		} else {

			body, err := cs.tika.Parse(context.Background(), f)

			if err != nil {
				e := fmt.Errorf("Error with TikaClient parse: %v", err)
				log.Fatal(e)
			} else {

				body := strip.StripTags(body)

				caseStudyObj := cs.saveCaseStudy(body, url, companyNumber)

				dto := web.TranslateToElastic(*caseStudyObj)

				ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

				esErr := cs.es.PutRecord(ctx, dto)

				if esErr != nil {
					log.Fatalf("failed to put record into elasticsearch: %v", err)
				}

				defer f.Close()
				defer os.Remove(fileName)
				enc.Encode(dto)
			}
		}
	}
}

func (cs *CaseStudyService) saveCaseStudy(body string, url string, companyNumber string) *model.CaseStudy {
	b := strings.TrimSpace(body)
	b = strings.Replace(b, "\n", "", -1)

	//Substring methodoloy
	runes := []rune(b)

	var caseStudyText string

	if len(runes) > 4500 {
		caseStudyText = string(runes[0:4500])
	} else {
		caseStudyText = string(runes)
	}

	str, _ := uuid.NewRandom()

	caseStudyObj := cs.dbs.DB.FindCaseStudyBySourceAndCompanyNumber(url, companyNumber)

	if caseStudyObj.ID == "" {
		caseStudyObj.ID = str.String()
	}

	caseStudyObj.Title = web.GetFileName(url)
	caseStudyObj.IdentifiedOn = time.Now()
	caseStudyObj.CaseStudyText = caseStudyText

	as, _ := aws.RunComprehend([]string{caseStudyObj.CaseStudyText})

	companies, people := aws.DetermineOrganisationTag(as)

	cs.dbs.DB.DeleteCaseStudyOrganisations(caseStudyObj.ID)

	companyArr := []model.CaseStudyOrganisations{}
	peopleArr := []model.CaseStudyPeople{}

	for _, o := range companies {
		test := cs.dbs.DB.FindCaseStudyOrganisationByNameAndCaseID(o, caseStudyObj.ID)
		obj := cs.dbs.DB.SaveCaseStudyOrganisation(test)
		companyArr = append(companyArr, *obj)
	}

	for _, o := range people {
		test := cs.dbs.DB.FindCaseStudyPersonByNameAndCaseID(o, caseStudyObj.ID)
		obj := cs.dbs.DB.SaveCaseStudyPerson(test)
		peopleArr = append(peopleArr, *obj)
	}

	caseStudyObj.Organizations = companyArr
	caseStudyObj.People = peopleArr

	return cs.dbs.DB.SaveCaseStudy(caseStudyObj)
}
