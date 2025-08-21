package models

type Course struct {
	CourseID      int    `json:"course_id"`
	Degree        string `json:"degree"`
	Major         string `json:"major"`
	Year          int    `json:"year"`
	ThaiCourse    string `json:"thai_course"`
	EngCourse     string `json:"eng_course"`
	ThaiDegree    string `json:"thai_degree"`
	EngDegree     string `json:"eng_degree"`
	AdmissionReq  string `json:"admission_req"`
	GraduationReq string `json:"graduation_req"`
	Philosophy    string `json:"philosophy"`
	Objective     string `json:"objective"`
	Tuition       string `json:"tuition"`
	Credits       string `json:"credits"`
	CareerPaths   string `json:"career_paths"`
	PLO           string `json:"plo"`
	DetailURL     string `json:"detail_url"`
}

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
