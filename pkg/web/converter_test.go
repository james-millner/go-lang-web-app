package web

import (
	"testing"
	"time"

	"github.com/james-millner/go-lang-web-app/pkg/model"
<<<<<<< HEAD

=======
>>>>>>> master
	"github.com/stretchr/testify/assert"
)

func TestTranslateToElastic(t *testing.T) {

	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	timeNow, _ := time.Parse(layout, str)

	organisation := &model.CaseStudyOrganisations{CaseStudyID: "1", OrganisationName: "IQBlade"}

	orgArray := []model.CaseStudyOrganisations{*organisation}

	caseStudyObj := &model.CaseStudy{ID: "1", Title: "Super Duper CaseStudy", CompanyNumber: "12345", CaseStudyText: "This is the text", SourceURL: "www.superdupercasestudy.com/case-studies/Super-Duper-CaseStudy.pdf?s=98234", Organizations: orgArray, CreatedAt: timeNow, UpdatedAt: timeNow, IdentifiedOn: timeNow}

	caseStudyDTO := TranslateToElastic(*caseStudyObj)

	assert.Equal(t, caseStudyObj.SourceURL, caseStudyDTO.SourceURL)
	assert.Equal(t, caseStudyObj.CaseStudyText, caseStudyDTO.CaseStudyText)
	assert.Equal(t, caseStudyObj.CompanyNumber, caseStudyDTO.CompanyNumber)
	assert.Equal(t, caseStudyDTO.CreatedAt, "2014-11-12")
	assert.Equal(t, caseStudyDTO.UpdatedAt, "2014-11-12")
	assert.Equal(t, caseStudyDTO.IdentifiedOn, "2014-11-12")
	assert.Equal(t, caseStudyObj.Title, caseStudyDTO.Title)
	assert.Len(t, caseStudyDTO.Organizations, 1)

}
