package model

import "time"

//Response Object
type Response struct {
	ID        uint   `gorm:"primary_key"`
	SourceURL string `gorm:"size:255;index:idx_name_response"`
	Success   bool
	URLFound  string `gorm:"index:idx_name_response"`
	Document  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//CaseStudy entity
type CaseStudy struct {
	ID            uint   `gorm:"primary_key"`
	CompanyNumber string `gorm:"index:idx_company_number" json:"companyNumber"`
	SourceURL     string `gorm:"index:idx_source_url" json:"sourceUrl"`
	CaseStudyText string `gorm:"size:500" json:"caseStudyText"`
	IdentifiedOn  time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
