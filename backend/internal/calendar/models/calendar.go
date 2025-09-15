package models

type Calendar struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Location    string `json:"location,omitempty"`
	HtmlLink    string `json:"htmlLink,omitempty"`
}
