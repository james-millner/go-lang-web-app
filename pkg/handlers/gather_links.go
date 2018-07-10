package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/james-millner/go-lang-web-app/pkg/service"
	"github.com/james-millner/go-lang-web-app/pkg/web"
)

//GatherInterface going to be used for testing.
type GatherInterface interface {
	GetLinks(url string) []string
	GatherLinks(url string) []string
	ProcessLinks(url string, results []string) []string
}

// ResponseService to be used to handle communication to the DB and Service Methods.
type ResponseService struct {
	rs *service.ResponseService
	gi GatherInterface
}

//NewResponseService constructor
func NewResponseService(rs *service.ResponseService) *ResponseService {
	return &ResponseService{
		rs: rs,
	}
}

//GatherLinks function used to process links for a given URL.
func (rs *ResponseService) GatherLinks() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		url := r.FormValue("url")

		if url == "" {
			resp := &model.Response{Success: false}
			enc.Encode(resp)
		}

		fmt.Println("Starting to gather all first and secondary links for: " + url)

		var array []string

		results := rs.ProcessLinks(url, array)

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

			document := rs.rs.DB.FindBySourceURLAndURLFound(url, result)
			document.Success = true
			document.Document = isDoc
			rs.rs.DB.Save(document)

		}
		resp := &model.ResponseDTO{Links: links, Documents: documents, SourceURL: url}

		fmt.Println("Total links found: ", len(results))
		fmt.Println("Total case study links found for:", len(links))
		fmt.Println("Total documents found for:", len(documents))

		fmt.Println("----------")

		for _, a := range links {
			fmt.Println(a)
		}

		fmt.Println("----------")

		for _, a := range documents {
			fmt.Println(a)
		}

		fmt.Println("Finished processing: " + url)

		enc.Encode(resp)
	}
}

//ProcessLinks method
func (rs *ResponseService) ProcessLinks(url string, results []string) []string {

	var processed []string

	//First Iteration of the given URL.
	for _, u := range rs.GetLinks(url) {
		if !web.ArrayContains(results, u) && web.IsProbableLink(u) {
			results = append(results, u)
		}
	}

	fmt.Println("First iteration complete.")
	fmt.Println(fmt.Sprintf("%s%d", "Results size: ", len(results)))

	for _, u := range results {
		for _, s := range rs.GetLinks(u) {
			if !web.ArrayContains(results, s) && web.IsProbableLink(s) {
				results = append(results, s)
			}
		}
		processed = append(processed, u)
	}

	fmt.Println("Second iteration complete.")
	fmt.Println(fmt.Sprintf("%s%d", "Results size: ", len(results)))

	for _, u := range results {
		if !web.ArrayContains(processed, u) {
			for _, s := range rs.GetLinks(u) {
				if !web.ArrayContains(results, s) && web.IsProbableLink(s) {
					results = append(results, s)
				}
			}
			processed = append(processed, u)
		}
	}

	fmt.Println("Third iteration complete.")
	fmt.Println(fmt.Sprintf("%s%d", "Results size: ", len(results)))

	return results
}

//GetLinks Method
func (rs *ResponseService) GetLinks(url string) []string {

	d, err := web.GetResponse(url)

	if err != nil {
		errFmt := fmt.Errorf("failed to execute request: %v", err)
		fmt.Println(errFmt)
		return nil
	}
	return web.RetreiveLinksFromDocument(d)
}
