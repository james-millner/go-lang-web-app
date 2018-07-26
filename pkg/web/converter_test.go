package web

import (
	"github.com/james-millner/go-lang-web-app/pkg/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTranslateToElastic(t *testing.T) {

	timeNow := time.Now()

	organisation := &model.CaseStudyOrganisations{CaseStudyID: "1", OrganisationName: "IQBlade"}

	orgArray := []model.CaseStudyOrganisations{*organisation}

	caseStudyObj := &model.CaseStudy{ID: "1", Title: "Super Duper CaseStudy", CompanyNumber: "12345", CaseStudyText: "This is the text", SourceURL: "www.superdupercasestudy.com/case-studies/Super-Duper-CaseStudy.pdf?s=98234", Organizations: orgArray, CreatedAt: timeNow, UpdatedAt: timeNow, IdentifiedOn: timeNow}

	caseStudyDTO := TranslateToElastic(*caseStudyObj)

	assert.Equal(t, caseStudyObj.SourceURL, caseStudyDTO.SourceURL)
	assert.Equal(t, caseStudyObj.CaseStudyText, caseStudyDTO.CaseStudyText)
	assert.Equal(t, caseStudyObj.CompanyNumber, caseStudyDTO.CompanyNumber)
	assert.Equal(t, caseStudyObj.CreatedAt, caseStudyDTO.CreatedAt)
	assert.Equal(t, caseStudyObj.UpdatedAt, caseStudyDTO.UpdatedAt)
	assert.Equal(t, caseStudyObj.IdentifiedOn, caseStudyDTO.IdentifiedOn)
	assert.Equal(t, caseStudyObj.Title, caseStudyDTO.Title)
	assert.Len(t, caseStudyDTO.Organizations, 1)

}
