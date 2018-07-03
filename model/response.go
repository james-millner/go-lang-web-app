package model

type Response struct {
	ID			uint
	Url 		string 		`gorm:"size:255;"`
	Success   	bool
	Links     	[]Links		`gorm:"many2many:response_links;"`
	Documents 	[]Links		`gorm:"many2many:response_documents;"`
}

type IndividualLinkResponse struct {
	Url			string
	Selector	string
	Tag			string
}

type Links struct {
	ID 		uint
	Url 	string
}
