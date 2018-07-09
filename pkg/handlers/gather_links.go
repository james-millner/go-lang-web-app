package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/james-millner/go-lang-web-app/pkg/service"
	"github.com/james-millner/go-lang-web-app/pkg/web"
	"github.com/jinzhu/gorm"
)

type GatherInterface interface {
	GatherLinks(url string) []string
	ProcessLinks(url string, results []string) []string
}

type Service struct {
	Storage *gorm.DB
	Router  *mux.Router
	debug   bool
}

// User is used to provide http handlers for retrieving Twitter User profiles and tweets
type ResponseTest struct {
	rs *service.ResponseService
}

//GatherLinks function used to process links for a given URL.
func (a *ResponseTest) GatherLinks() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		url := r.FormValue("url")

		if url == "" {
			resp := &model.Response{Success: false}
			enc.Encode(resp)
		}

		fmt.Println("Starting to gather all first and secondary links.")

		var array []string

		results := a.ProcessLinks(url, array)

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

			document := a.rs.DB.FindBySourceURLAndURLFound(url, result)
			document.Success = true
			document.Document = isDoc
			a.rs.DB.Save(document)

		}
		resp := &model.ResponseDTO{Links: links, Documents: documents, SourceURL: url}

		fmt.Println("Total links found: ", len(results))
		fmt.Println("Total case study links found for:", len(links))
		fmt.Println("Total documents found for:", len(documents))

		for _, a := range links {
			fmt.Println(a)
		}

		fmt.Println("----------")

		for _, a := range documents {
			fmt.Println(a)
		}

		enc.Encode(resp)
	}
}

//ProcessLinks
func (r *ResponseTest) ProcessLinks(url string, results []string) []string {
	for _, u := range web.GetLinks(url) {
		if !web.ArrayContains(results, u) && web.IsProbableLink(u) {
			results = append(results, u)
		}
	}

	for _, u := range results {
		for _, s := range web.GetLinks(u) {
			if !web.ArrayContains(results, s) && web.IsProbableLink(s) {
				results = append(results, s)
			}
		}
	}

	return results
}
