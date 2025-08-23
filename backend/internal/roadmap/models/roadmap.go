package models

type Roadmap struct {
	RoadmapID  int    `json:"roadmap_id"`
	CourseID   int    `json:"course_id"`
	ThaiCourse string `json:"thai_course"`
	RoadmapURL string `json:"roadmap_url"`
}

type RoadmapQueryParam struct {
	Search   string `form:"search"`
	Limit    int    `form:"limit"`
	CourseID int    `form:"course_id"`
	Sort     string `form:"sort"`
	Order    string `form:"order"`
}

type RoadmapRequest struct {
	CourseID   int    `json:"course_id"`
	RoadmapURL string `json:"roadmap_url"`
}
