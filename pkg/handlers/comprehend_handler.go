package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/james-millner/go-lang-web-app/pkg/service"
)

// ResponseService to be used to handle communication to the DB and Service Methods.
type CaseStudyService struct {
	dbs *service.DBService
}

//ComprehendCaseStudy function.
func (rs *CaseStudyService) ComprehendCaseStudy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		id := r.FormValue("id")
		log.Println(id)

	}
}
