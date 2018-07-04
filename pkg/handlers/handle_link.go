package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/james-millner/go-lang-web-app/pkg/model"
	"github.com/jinzhu/gorm"
)

//HandleLinks Function for processing individual links
func HandleLinks(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		selector := r.FormValue("selector")

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		resp := &model.ProcessLinkDTO{SourceURL: url, Selector: selector}
		enc.Encode(resp)
	}
}
