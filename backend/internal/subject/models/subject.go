package models

type Subjects struct {
	ID                int     `json:"id"`
	SubjectID         string  `json:"subject_id"`
	CourseID          int     `json:"course_id"`
	ThaiCourse        string  `json:"thai_course"`
	PlanType          string  `json:"plan_type"`
	Semester          string  `json:"semester"`
	ThaiSubject       string  `json:"thai_subject"`
	EngSubject        *string `json:"eng_subject"`
	Credits           string  `json:"credits"`
	CompulsorySubject *string `json:"compulsory_subject"`
	Condition         *string `json:"condition"`
	DescriptionThai   *string `json:"description_thai"`
	DescriptionEng    *string `json:"description_eng"`
	CLO               *string `json:"clo"`
}

type SubjectsQueryParam struct {
	Search    string `form:"search"`
	Limit     int    `form:"limit"`
	SubjectID string `form:"subject_id"`
	CourseID  int    `form:"course_id"`
	PlanType  string `form:"plan_type"`
	Semester  string `form:"semester"`
	Sort      string `form:"sort"`
	Order     string `form:"order"`
}

type SubjectsRequest struct {
	ID                int     `json:"id"`
	SubjectID         string  `json:"subject_id"`
	CourseID          int     `json:"course_id"`
	PlanType          string  `json:"plan_type"`
	Semester          string  `json:"semester"`
	ThaiSubject       string  `json:"thai_subject"`
	EngSubject        *string `json:"eng_subject"`
	Credits           string  `json:"credits"`
	CompulsorySubject *string `json:"compulsory_subject"`
	Condition         *string `json:"condition"`
	DescriptionThai   *string `json:"description_thai"`
	DescriptionEng    *string `json:"description_eng"`
	CLO               *string `json:"clo"`
}
