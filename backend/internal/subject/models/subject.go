package models

type Subjects struct {
	ID                int     `json:"id"`
	SubjectID         string  `json:"subject_id"`
	CourseID          string  `json:"course_id"`
	ThaiCourse        string  `json:"thai_course"`
	PlanTypeID        int     `json:"plan_type_id"`
	PlanType          string  `json:"plan_type"`
	SemesterID        int     `json:"semester_id"`
	Semester          string  `json:"semester"`
	ThaiSubject       string  `json:"thai_subject"`
	EngSubject        *string `json:"eng_subject"`
	Credits           string  `json:"credits"`
	CompulsorySubject *string `json:"compulsory_subject"`
	Condition         *string `json:"condition"`
	DescriptionID     *string `json:"description_id"`
	DescriptionThai   *string `json:"description_thai"`
	DescriptionEng    *string `json:"description_eng"`
	CloID             *string `json:"clo_id"`
	CLO               *string `json:"clo"`
}

type SubjectsQueryParam struct {
	Search     string `form:"search"`
	Limit      int    `form:"limit"`
	SubjectID  string `form:"subject_id"`
	CourseID   int    `form:"course_id"`
	PlanTypeID int    `form:"plan_type_id"`
	SemesterID int    `form:"semester_id"`
	Sort       string `form:"sort"`
	Order      string `form:"order"`
}

type SubjectsRequest struct {
	ID                int     `json:"id"`
	SubjectID         string  `json:"subject_id"`
	CourseID          string  `json:"course_id"`
	PlanTypeID        int     `json:"plan_type_id"`
	SemesterID        int     `json:"semester_id"`
	ThaiSubject       string  `json:"thai_subject"`
	EngSubject        *string `json:"eng_subject"`
	Credits           string  `json:"credits"`
	CompulsorySubject *string `json:"compulsory_subject"`
	Condition         *string `json:"condition"`
	DescriptionID     *string `json:"description_id"`
	DescriptionThai   *string `json:"description_thai"`
	DescriptionEng    *string `json:"description_eng"`
	CloID             *string `json:"clo_id"`
	CLO               *string `json:"clo"`
}
