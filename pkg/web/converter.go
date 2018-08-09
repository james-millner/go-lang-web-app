package web

import "github.com/iqblade/casestudies/pkg/model"

//TranslateToElastic method converts a model.CaseStudy object into a DTO to be returned as a response body.
func TranslateToElastic(obj model.CaseStudy) model.CaseStudyDTO {
	var dto model.CaseStudyDTO

	dto.ID = obj.ID
	dto.Title = obj.Title
	dto.CompanyNumber = obj.CompanyNumber
	dto.CaseStudyText = obj.CaseStudyText
	dto.SourceURL = obj.SourceURL
	dto.CreatedAt = obj.CreatedAt.Format("2006-01-02")
	dto.IdentifiedOn = obj.IdentifiedOn.Format("2006-01-02")
	dto.UpdatedAt = obj.UpdatedAt.Format("2006-01-02")

	companies := []string{}
	people := []string{}
	locations := []string{}

	for _, o := range obj.Organizations {
		companies = append(companies, o.OrganisationName)
	}

	for _, p := range obj.People {
		people = append(people, p.PersonName)
	}

	for _, l := range obj.Locations {
		locations = append(locations, l.Location)
	}

	dto.Organizations = companies
	dto.People = people
	dto.Locations = locations

	return dto

}
