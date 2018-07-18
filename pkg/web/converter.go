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

	companies := []string{}
	people := []string{}

	for _, o := range obj.Organizations {
		companies = append(companies, o.OrganisationName)
	}

	for _, p := range obj.People {
		people = append(people, p.PersonName)
	}

	dto.Organizations = companies
	dto.People = people

	return dto

}
