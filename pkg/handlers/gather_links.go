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

		for _, firstLevel := range getLinks(url) {
			//Iterate over all the immediet links

			if !web.ArrayContains(results, firstLevel) {
				results = append(results, firstLevel)
			}

			for _, secondLevel := range getLinks(firstLevel) {
				if !web.ArrayContains(results, secondLevel) {
					results = append(results, secondLevel)
				}
			}
		}

		var links []string
		var documents []string

		for _, result := range results {
			if web.IsPDFDocument(result) {
				documents = append(documents, result)
			} else {
				links = append(links, result)
			}
		}

		resp := &model.ResponseDTO{Links: links, Documents: documents, SourceURL: url}

		fmt.Println("Total links found for:", len(links))
		fmt.Println("Total documents found for:", len(documents))

		for l := range results {
			fmt.Println(links[l])
		}

		enc.Encode(resp)
	}
}

func getLinks(url string) []string {

	fmt.Println(url)

	r, err := web.GetResponse(url)

	if err != nil {
		errFmt := fmt.Errorf("failed to execute request: %v", err)
		fmt.Println(errFmt)
		return nil
	}

	return web.RetreiveLinksFromDocument(r)
}
