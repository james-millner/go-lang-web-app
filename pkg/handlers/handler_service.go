package handlers

import (
	"github.com/google/go-tika/tika"
	"github.com/james-millner/go-lang-web-app/pkg/es"
	"github.com/james-millner/go-lang-web-app/pkg/service"
)

// ResponseService to be used to handle communication to the DB and Service Methods.
type CaseStudyService struct {
	dbs  *service.CaseStudyService
	tika *tika.Client
	es   *es.Elastic
}

//NewCaseStudyService constructor
func NewCaseStudyService(dbs *service.CaseStudyService, tc *tika.Client, es *es.Elastic) *CaseStudyService {
	return &CaseStudyService{
		dbs:  dbs,
		tika: tc,
		es:   es,
	}
}
