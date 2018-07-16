package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

//ComprehendCaseStudy function.
func (cs *CaseStudyService) ComprehendCaseStudy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		id := r.FormValue("id")
		log.Println(id)

	}
}
