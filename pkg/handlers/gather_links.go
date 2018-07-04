package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/james-millner/go-lang-web-app/pkg/web"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Storage *gorm.DB
	Router  *mux.Router
	debug   bool
}

func GatherLinks(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		url := r.FormValue("url")

		if url != "" {

			var links []string
			var documents []string

			for _, t := range getLinks(url) {
				if strings.Contains(t, ".pdf") {
					documents = append(documents, t)
				} else {
					links = append(links, t)
				}
			}

			resp := &model.ResponseDTO{Links: links, Documents: documents, SourceURL: url}

			for l := range links {
				fmt.Println(links[l])
			}

			fmt.Println("Total links found for:", len(links))
			fmt.Println("Total documents found for:", len(documents))

			enc.Encode(resp)

		} else {
			resp := &model.Response{Success: false}
			enc.Encode(resp)
		}
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

	return web.GetPageLinks(r)
}
