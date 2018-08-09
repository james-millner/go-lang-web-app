package handlers

import (
	"github.com/iqblade/casestudies/pkg/es"
	"github.com/iqblade/casestudies/pkg/service"

	"github.com/google/go-tika/tika"
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
