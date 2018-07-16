package web

import "github.com/james-millner/go-lang-web-app/pkg/model"

func TranslateToElastic(obj model.CaseStudy) model.CaseStudyDTO {
	var dto model.CaseStudyDTO

	dto.ID = obj.ID
	dto.CompanyNumber = obj.CompanyNumber
	dto.CaseStudyText = obj.CaseStudyText
	dto.SourceURL = obj.SourceURL
	dto.CreatedAt = obj.CreatedAt
	dto.IdentifiedOn = obj.IdentifiedOn
	dto.UpdatedAt = obj.UpdatedAt

	strArray := []string{}

	for _, o := range obj.Organizations {
		strArray = append(strArray, o.OrganisationName)
	}

	dto.Organizations = strArray

	return dto

}
