package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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

		if url != "" {

			var links []string
			var documents []string

			for _, t := range getLinks(url) {
				if strings.Contains(t, ".pdf") {
					documents = append(documents, t)
					a.rs.DB.Save(&model.Response{SourceURL: url, URLFound: t, CreatedAt: time.Now(), Success: true, DocumentType: 0})

				} else {
					links = append(links, t)
					a.rs.DB.Save(&model.Response{SourceURL: url, URLFound: t, CreatedAt: time.Now(), Success: true, DocumentType: 1})

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
