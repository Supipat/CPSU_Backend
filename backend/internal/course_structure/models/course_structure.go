package models

type CourseStructure struct {
	CourseStructureID  int    `json:"course_structure_id"`
	CourseID           int    `json:"course_id"`
	ThaiCourse         string `json:"thai_course"`
	CourseStructureURL string `json:"course_structure_url"`
}

type CourseStructureQueryParam struct {
	Search   string `form:"search"`
	Limit    int    `form:"limit"`
	CourseID int    `form:"course_id"`
	Sort     string `form:"sort"`
	Order    string `form:"order"`
}

type CourseStructureRequest struct {
	CourseID           int    `json:"course_id"`
	CourseStructureURL string `json:"course_structure_url"`
}
