package model

import "time"

//Response Object
type Response struct {
	ID        uint   `gorm:"primary_key"`
	SourceURL string `gorm:"size:200;" index:"idx_name_response"`
	Success   bool
	URLFound  string `gorm:"index:idx_name_response"`
	Document  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//CaseStudy entity
type CaseStudy struct {
	ID            string                   `gorm:"primary_key" size:"70" json:"id"`
	CompanyNumber string                   `gorm:"index:idx_company_number" json:"companyNumber" size:"20"`
	SourceURL     string                   `gorm:"index:idx_source_url" json:"sourceUrl" size:"200"`
	CaseStudyText string                   `gorm:"size:7500" json:"caseStudyText"`
	Organizations []CaseStudyOrganisations `gorm:"one2many:case_studies_organisations;" json:"organisations"`
	People        []CaseStudyPeople        `gorm:"one2many:case_studies_people;" json:"people"`
	IdentifiedOn  time.Time                `json:"identifiedOn"`
	CreatedAt     time.Time                `json:"createdAt"`
	UpdatedAt     time.Time                `json:"updatedAt"`
}

//CaseStudyOrganisations entity
type CaseStudyOrganisations struct {
	CaseStudyID      string `json:"caseStudyId"`
	OrganisationName string `gorm:"index:idx_organisation" json:"organisationName"`
}

//CaseStudyPeople entity
type CaseStudyPeople struct {
	CaseStudyID string `json:"caseStudyId"`
	PersonName  string `gorm:"index:idx_person_name" json:"peopleName"`
}
