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

				runes := []rune(body)
				//... Convert back into a string from rune slice.
				safeSubstring := string(runes[0:4500])

				csss := cs.dbs.DB.FindCaseStudyBySourceAndCompanyNumber(url, companyNumber)
				csss.IdentifiedOn = time.Now()
				csss.CaseStudyText = safeSubstring

				saved := cs.dbs.DB.SaveCaseStudy(csss)

				as, _ := aws.RunComprehend([]string{safeSubstring})

				companies := aws.DetermineOrganisationTag(as)

				log.Println(fmt.Sprintf("%v%v", len(companies), " companies found!"))

				cs.dbs.DB.DeleteCaseStudyOrganisations(saved.ID)

				for _, o := range companies {
					test := cs.dbs.DB.FindCaseStudyOrganisationByNameAndCaseID(o, saved.ID)
					cs.dbs.DB.SaveCaseStudyOrganisation(test)
					log.Println("Saved..")
				}

				f.Close()
				os.Remove(fileName)
				//log.Println(body)
			}
		}
	}
}
