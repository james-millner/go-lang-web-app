package model

import "time"

//ResponseDTO Object for returing to the user / client.
type ResponseDTO struct {
	SourceURL string   `json:"source"`
	Links     []string `json:"link"`
	Documents []string `json:"documents"`
}

//CaseStudy entity
type CaseStudyDTO struct {
	ID            string
	CompanyNumber string    `json:"companyNumber"`
	SourceURL     string    `json:"sourceUrl"`
	CaseStudyText string    `json:"caseStudyText"`
	Organizations []string  `json:"organisations"`
	People        []string  `json:"people"`
	IdentifiedOn  time.Time `json:"identifiedOn"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
