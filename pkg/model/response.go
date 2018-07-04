package model

type Response struct {
	ID			uint
	SourceURL 	string 		`gorm:"size:255;"`
	Success   	bool
	URLFound    string		`gorm:"many2many:response_links;"`
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

type ResponseDTO struct{
	SourceURL 	string 		`json:"source"`
	Links    	[]string	`json:"link"`
	Documents 	[]string	`json:"documents"`
}


type EnumValue int

const (
	DOCUMENT EnumValue = 0
	HTML_LINK EnumValue = 1
)
