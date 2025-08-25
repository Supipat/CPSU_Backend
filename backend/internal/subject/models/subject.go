package models

type Subjects struct {
	SubjectID         int    `json:"subject_id"`
	CourseID          int    `json:"course_id"`
	Degree            string `json:"degree"`
	Major             string `json:"major"`
	Year              int    `json:"year"`
	PlanType          string `json:"plan_type"`
	Semester          string `json:"semester"`
	ThaiSubject       string `json:"thai_subject"`
	EngSubject        string `json:"eng_subject"`
	Credits           string `json:"credits"`
	CompulsorySubject string `json:"compulsory_subject"`
	Condition         string `json:"condition"`
	DescriptionThai   string `json:"description_thai"`
	DescriptionEng    string `json:"description_eng"`
	CLO               string `json:"clo"`
}

type SubjectsQueryParam struct {
	Search    string `form:"search"`
	Limit     int    `form:"limit"`
	SubjectID int    `json:"subject_id"`
	CourseID  int    `json:"course_id"`
	PlanType  string `json:"plan_type"`
	Semester  string `json:"semester"`
	Sort      string `form:"sort"`
	Order     string `form:"order"`
}

type SubjectsRequest struct {
	CourseID          int    `json:"course_id"`
	PlanType          string `json:"plan_type"`
	Semester          string `json:"semester"`
	ThaiSubject       string `json:"thai_subject"`
	EngSubject        string `json:"eng_subject"`
	Credits           string `json:"credits"`
	CompulsorySubject string `json:"compulsory_subject"`
	Condition         string `json:"condition"`
	DescriptionThai   string `json:"description_thai"`
	DescriptionEng    string `json:"description_eng"`
	CLO               string `json:"clo"`
}
