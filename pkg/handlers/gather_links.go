package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/james-millner/go-lang-web-app/pkg/web"
)

//GatherLinks function used to process links for a given URL.
func (cs *CaseStudyService) GatherLinks() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		log.Println(r)

		url := r.FormValue("url")

		if url == "" {
			resp := &model.Response{Success: false}
			enc.Encode(resp)
			return
		}

		fmt.Println("Starting to gather all first and secondary links for: " + url)

		var array []string

		results := cs.GatherLinksFromBaseURL(url, array)

		fmt.Println("Finished Looping.")

		var links []string
		var documents []string

		fmt.Println("Results found: ", len(results))

		for _, result := range results {

			isDoc := web.IsPDFDocument(result)

			if isDoc {
				documents = append(documents, result)
			} else {
				links = append(links, result)
			}

			document := cs.dbs.DB.FindBySourceURLAndURLFound(url, result)
			document.Success = true
			document.Document = isDoc
			cs.dbs.DB.SaveResponse(document)

		}
		resp := &model.ResponseDTO{Links: links, Documents: documents, SourceURL: url}

		fmt.Println("Total links found: ", len(results))
		fmt.Println("Total case study links found for:", len(links))
		fmt.Println("Total documents found for:", len(documents))
		fmt.Println("----------")
		fmt.Println("Finished processing: " + url)

		enc.Encode(resp)
	}
}

//GatherLinksFromBaseURL method
func (cs *CaseStudyService) GatherLinksFromBaseURL(url string, results []string) []string {

	var processed []string

	initialLinks, _ := cs.GetLinks(url)

	//First Iteration of the given URL.
	for _, u := range initialLinks {
		up := validateURLStructure(url, u)

		log.Println(up)

		if !web.ArrayContains(results, up) && web.IsPossibleCaseStudyLink(up) {
			results = append(results, up)
		}
	}

	processed = append(processed, url)

	fmt.Println("First iteration complete.")
	fmt.Println(fmt.Sprintf("%s%d", "Results size: ", len(results)))

	for _, u := range results {
		if !web.ArrayContains(processed, u) {
			secondaryLinks, _ := cs.GetLinks(u)
			for _, s := range secondaryLinks {
				up := validateURLStructure(url, s)

				log.Println(up)

				processed, results = update(up, processed, results)
			}
			processed = append(processed, u)
		}
	}

	fmt.Println("Second iteration complete.")
	fmt.Println(fmt.Sprintf("%s%d", "Results size: ", len(results)))

	for _, u := range results {
		if !web.ArrayContains(processed, u) {
			thirdLevelLinks, _ := cs.GetLinks(u)
			for _, s := range thirdLevelLinks {
				up := validateURLStructure(url, s)

				log.Println(up)

				processed, results = update(up, processed, results)
			}
			processed = append(processed, u)
		}
	}

	fmt.Println("Third iteration complete.")
	fmt.Println(fmt.Sprintf("%s%d", "Results size: ", len(results)))
	log.Println("Finished gathering..")

	return results
}

func validateURLStructure(baseurl string, foundurl string) string {
	if !strings.Contains(foundurl, "http") {
		return baseurl + foundurl
	}

	return foundurl
}

func update(url string, processed []string, results []string) ([]string, []string) {
	if !web.ArrayContains(processed, url) {
		if !web.ArrayContains(results, url) && web.IsPossibleCaseStudyLink(url) {
			results = append(results, url)
		}
	}

	return processed, results
}

//GetLinks Method
func (cs *CaseStudyService) GetLinks(url string) ([]string, error) {

	resp, err := http.Get(url)

	if err != nil {
		e := fmt.Errorf("failed to execute request: %v", err)
		fmt.Println(e)
		return []string{}, err
	}

	defer resp.Body.Close()

	if err != nil {
		errFmt := fmt.Errorf("failed to execute request: %v", err)
		fmt.Println(errFmt)
		return []string{}, errFmt
	}

	doc, err := web.GetGoqueryDocument(resp.Body)

	return web.RetreiveLinksFromDocument(url, doc), err
}
