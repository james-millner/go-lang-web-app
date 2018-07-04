package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/james-millner/go-lang-web-app/model"
	"github.com/jinzhu/gorm"
)

func HandleLinks(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		selector := r.FormValue("selector")
		tag := r.FormValue("tag")

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		resp := &model.IndividualLinkResponse{Url: url, Selector: selector, Tag: tag}
		enc.Encode(resp)
	}
}
