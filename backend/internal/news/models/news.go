package models

import "time"

type News struct {
	NewsID    int          `json:"news_id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	TypeID    int          `json:"type_id"`
	TypeName  string       `json:"type_name"`
	DetailURL string       `json:"detail_url"`
	Images    []NewsImages `json:"images"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"update_at"`
}

type NewsImages struct {
	ImageID   int    `json:"image_id"`
	NewsID    int    `json:"news_id"`
	FileImage string `json:"file_image"`
}

type NewsTypes struct {
	TypeID   string `json:"type_id"`
	NewsID   int    `json:"news_id"`
	TypeName string `json:"type_name"`
}

type NewsQueryParam struct {
	Search   string `form:"search"`
	Limit    int    `form:"limit"`
	TypeID   int    `form:"type_id"`
	TypeName string `form:"type_name"`
	Sort     string `form:"sort"`
	Order    string `form:"order"`
}

type NewsRequest struct {
	NewsID    int          `json:"news_id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	TypeID    int          `json:"type_id"`
	DetailURL string       `json:"detail_url"`
	Images    []NewsImages `json:"images"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"update_at"`
}
