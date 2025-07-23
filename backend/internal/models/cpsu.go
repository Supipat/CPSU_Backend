package models

import "time"

type News struct {
	NewsID    int         `json:"news_id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	NewsType  string      `json:"news_type"`
	DetailURL string      `json:"detail_url"`
	CreatedAt time.Time   `json:"create_at"`
	UpdatedAt time.Time   `json:"update_at"`
	Images    []NewsImage `json:"images"`
}

type NewsImage struct {
	ImageID  int    `json:"image_id"`
	NewsID   int    `json:"news_id"`
	ImageURL string `json:"image_url"`
}
