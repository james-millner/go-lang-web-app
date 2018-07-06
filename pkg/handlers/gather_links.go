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

type Service struct {
	Storage *gorm.DB
	Router  *mux.Router
	debug   bool
}

// User is used to provide http handlers for retrieving Twitter User profiles and tweets
type ResponseTest struct {
	rs *service.ResponseService
}

// NewUser creates a new Account struct that provides http handlers for Twitter Profile and Tweets
func NewUser(rs *service.ResponseService) *ResponseTest {
	return &ResponseTest{
		rs: rs,
	}
}

func (a *ResponseTest) GatherLinks() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		url := r.FormValue("url")

		if url == "" {
			resp := &model.Response{Success: false}
			enc.Encode(resp)
		}

		var results []string

		fmt.Println("Starting to gather all first and secondary links.")

		for _, firstLevel := range getLinks(url) {
			//Iterate over all the initial links from a given URL.
			if !web.ArrayContains(results, firstLevel) && web.IsProbableLink(firstLevel) {
				results = append(results, firstLevel)
			}

			for _, secondLevel := range getLinks(firstLevel) {
				if !web.ArrayContains(results, secondLevel) && web.IsProbableLink(secondLevel) {
					results = append(results, secondLevel)
				}
			}
		}

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

		enc.Encode(resp)
	}
}

func getLinks(url string) []string {

	r, err := web.GetResponse(url)

	if err != nil {
		errFmt := fmt.Errorf("failed to execute request: %v", err)
		fmt.Println(errFmt)
		return nil
	}

	return web.RetreiveLinksFromDocument(r)
}
