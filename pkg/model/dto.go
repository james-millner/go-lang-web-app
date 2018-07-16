package model

//ResponseDTO Object for returing to the user / client.
type ResponseDTO struct {
	SourceURL string   `json:"source"`
	Links     []string `json:"link"`
	Documents []string `json:"documents"`
}
