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

	"github.com/google/uuid"
	"github.com/james-millner/go-lang-web-app/pkg/aws"
	"github.com/james-millner/go-lang-web-app/pkg/model"
)

//ProcessCaseStudyLink function.
func (cs *CaseStudyService) ProcessCaseStudyLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		companyNumber := r.FormValue("company_number")
		url := r.FormValue("url")
		log.Println(companyNumber + " - " + url)

		tokens := strings.Split(url, "/")
		fileName := tokens[len(tokens)-1]

		out, oserr := os.Create(fileName)

		if oserr != nil {
			e := fmt.Errorf("Error with creating OS file: %v", oserr)
			log.Fatal(e)
		}

		resp, err := http.Get(url)
		if err != nil {
			e := fmt.Errorf("Error with GET request: %v", err)
			log.Fatal(e)
		}

		defer resp.Body.Close()

		io.Copy(out, resp.Body)

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		} else {

			body, err := cs.tika.Parse(context.Background(), f)

			if err != nil {
				e := fmt.Errorf("Error with TikaClient parse: %v", err)
				log.Fatal(e)
			} else {

				body := strings.TrimSpace(body)

				//Substring methodoloy
				runes := []rune(body)
				safeSubstring := string(runes[0:4500])

				str, _ := uuid.NewRandom()

				csss := cs.dbs.DB.FindCaseStudyBySourceAndCompanyNumber(url, companyNumber)

				if csss.ID == "" {
					csss.ID = str.String()
				}

				csss.IdentifiedOn = time.Now()
				csss.CaseStudyText = safeSubstring
				saved := cs.dbs.DB.SaveCaseStudy(csss)

				as, _ := aws.RunComprehend([]string{safeSubstring})

				companies := aws.DetermineOrganisationTag(as)

				cs.dbs.DB.DeleteCaseStudyOrganisations(saved.ID)

				arr := []model.CaseStudyOrganisations{}

				for _, o := range companies {
					test := cs.dbs.DB.FindCaseStudyOrganisationByNameAndCaseID(o, saved.ID)
					obj := cs.dbs.DB.SaveCaseStudyOrganisation(test)
					arr = append(arr, *obj)
				}

				csss.Organizations = arr

				ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

				esErr := cs.es.PutRecord(ctx, *csss)

				if esErr != nil {
					log.Fatalf("failed to put record into elasticsearch: %v", err)
				}

				f.Close()
				os.Remove(fileName)
				enc.Encode(csss)
			}
		}
	}
}
