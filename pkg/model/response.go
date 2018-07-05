package model

import "time"

//Response Object
type Response struct {
	ID           uint   `gorm:"primary_key"`
	SourceURL    string `gorm:"size:255;index:idx_name_response"`
	Success      bool
	URLFound     string `gorm:"index:idx_name_response"`
	DocumentType int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//ResponseDTO Object for returing to the user.
type ResponseDTO struct {
	SourceURL string   `json:"source"`
	Links     []string `json:"link"`
	Documents []string `json:"documents"`
}

//ProcessLinkDTO Object as dummy object for now. Likely to be removed / refactored.
type ProcessLinkDTO struct {
	SourceURL string
	Selector  string
}

// type EnumValue int

// const (
// 	DOCUMENT EnumValue = 0
// 	HTML_LINK EnumValue = 1
// )
