package model

//ResponseDTO Object for returing to the user / client.
type ResponseDTO struct {
	SourceURL string   `json:"source"`
	Links     []string `json:"link"`
	Documents []string `json:"documents"`
}

//CaseStudy entity
type CaseStudyDTO struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	CompanyNumber string   `json:"companyNumber"`
	SourceURL     string   `json:"sourceUrl"`
	CaseStudyText string   `json:"caseStudyText"`
	Organizations []string `json:"organisations"`
	People        []string `json:"people"`
	IdentifiedOn  string   `json:"identifiedOn"`
	CreatedAt     string   `json:"createdAt"`
	UpdatedAt     string   `json:"updatedAt"`
}
