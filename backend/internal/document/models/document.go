package models

type Document struct {
	DocumentID  int     `json:"document_id"`
	TypeID      int     `json:"type_id"`
	TypeName    string  `json:"type_name"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	File        string  `json:"file"`
}

type DocumentQueryParam struct {
	Search string `form:"search"`
	Limit  int    `form:"limit"`
	TypeID int    `form:"type_id"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`
}

type DocumentRequest struct {
	TypeID      int     `json:"type_id"`
	TypeName    string  `json:"type_name"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	File        string  `json:"file"`
}
