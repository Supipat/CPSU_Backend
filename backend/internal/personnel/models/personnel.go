package models

type Personnels struct {
	PersonnelID            int     `json:"personnel_id"`
	TypePersonnel          string  `json:"type_personnel"`
	DepartmentPositionID   int     `json:"department_position_id"`
	DepartmentPositionName string  `json:"department_position_name"`
	AcademicPositionID     *int    `json:"academic_position_id"`
	ThaiAcademicPosition   *string `json:"thai_academic_position"`
	EngAcademicPosition    *string `json:"eng_academic_position"`
	ThaiName               string  `json:"thai_name"`
	EngName                string  `json:"eng_name"`
	Education              *string `json:"education"`
	RelatedFields          *string `json:"related_fields"`
	Email                  *string `json:"email"`
	Website                *string `json:"website"`
	FileImage              string  `json:"file_image"`
}

type PersonnelQueryParam struct {
	Search               string `form:"search"`
	Limit                int    `form:"limit"`
	TypePersonnel        string `from:"type_personnel"`
	DepartmentPositionID int    `from:"department_position_id"`
	AcademicPositionID   *int   `from:"academic_position_id"`
	Sort                 string `form:"sort"`
	Order                string `form:"order"`
}

type PersonnelRequest struct {
	PersonnelID          int     `json:"personnel_id"`
	TypePersonnel        string  `json:"type_personnel"`
	DepartmentPositionID int     `json:"department_position_id"`
	AcademicPositionID   *int    `json:"academic_position_id"`
	ThaiAcademicPosition *string `json:"thai_academic_position"`
	EngAcademicPosition  *string `json:"eng_academic_position"`
	ThaiName             string  `json:"thai_name"`
	EngName              string  `json:"eng_name"`
	Education            *string `json:"education"`
	RelatedFields        *string `json:"related_fields"`
	Email                *string `json:"email"`
	Website              *string `json:"website"`
	FileImage            string  `json:"file_image"`
}
