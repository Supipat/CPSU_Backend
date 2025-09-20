package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"cpsu/internal/personnel/models"
)

type PersonnelRepository interface {
	GetAllPersonnels(param models.PersonnelQueryParam) ([]models.Personnels, error)
	GetPersonnelByID(id int) (*models.Personnels, error)
	CreatePersonnel(req models.PersonnelRequest) (*models.Personnels, error)
	UpdatePersonnel(id int, req models.PersonnelRequest) (*models.Personnels, error)
	DeletePersonnel(id int) error
}

type personnelRepository struct {
	db *sql.DB
}

func NewPersonnelRepository(db *sql.DB) PersonnelRepository {
	return &personnelRepository{db: db}
}

func (r *personnelRepository) GetAllPersonnels(param models.PersonnelQueryParam) ([]models.Personnels, error) {
	query := `
		SELECT
			p.personnel_id, p.type_personnel, d.department_position_id, d.department_position_name,
			a.academic_position_id, a.thai_academic_position, a.eng_academic_position, p.thai_name, 
			p.eng_name, p.education, p.related_fields, p.email, p.website, p.file_image
		FROM personnels p
		LEFT JOIN department_position d ON p.department_position_id = d.department_position_id
		LEFT JOIN academic_position a ON p.academic_position_id = a.academic_position_id
	`

	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if param.TypePersonnel != "" {
		conditions = append(conditions, "p.type_personnel = $"+strconv.Itoa(argIndex))
		args = append(args, param.TypePersonnel)
		argIndex++
	}

	if param.DepartmentPositionID > 0 {
		conditions = append(conditions, "d.department_position_id = $"+strconv.Itoa(argIndex))
		args = append(args, param.DepartmentPositionID)
		argIndex++
	}

	if param.AcademicPositionID != nil {
		conditions = append(conditions, "a.academic_position_id = $"+strconv.Itoa(argIndex))
		args = append(args, *param.AcademicPositionID)
		argIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	sort := "p.personnel_id"
	if param.Sort != "" {
		sort = "p." + param.Sort
	}

	order := "ASC"
	if strings.ToUpper(param.Order) == "DESC" {
		order = "DESC"
	}

	query += " ORDER BY " + sort + " " + order

	if param.Limit > 0 {
		query += " LIMIT " + strconv.Itoa(param.Limit)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var personnels []models.Personnels
	for rows.Next() {
		var personnel models.Personnels
		err := rows.Scan(
			&personnel.PersonnelID, &personnel.TypePersonnel, &personnel.DepartmentPositionID, &personnel.DepartmentPositionName,
			&personnel.AcademicPositionID, &personnel.ThaiAcademicPosition, &personnel.EngAcademicPosition,
			&personnel.ThaiName, &personnel.EngName, &personnel.Education, &personnel.RelatedFields,
			&personnel.Email, &personnel.Website, &personnel.FileImage,
		)
		if err != nil {
			return nil, err
		}
		personnels = append(personnels, personnel)
	}
	return personnels, nil
}

func (r *personnelRepository) GetPersonnelByID(id int) (*models.Personnels, error) {
	query := `
		SELECT
			p.personnel_id, p.type_personnel, d.department_position_id, d.department_position_name,
			a.academic_position_id, a.thai_academic_position, a.eng_academic_position, p.thai_name, 
			p.eng_name, p.education, p.related_fields, p.email, p.website, p.file_image
		FROM personnels p
		LEFT JOIN department_position d ON p.department_position_id = d.department_position_id
		LEFT JOIN academic_position a ON p.academic_position_id = a.academic_position_id
		WHERE p.personnel_id = $1
	`

	row := r.db.QueryRow(query, id)

	var personnel models.Personnels
	err := row.Scan(
		&personnel.PersonnelID, &personnel.TypePersonnel, &personnel.DepartmentPositionID, &personnel.DepartmentPositionName,
		&personnel.AcademicPositionID, &personnel.ThaiAcademicPosition, &personnel.EngAcademicPosition,
		&personnel.ThaiName, &personnel.EngName, &personnel.Education, &personnel.RelatedFields,
		&personnel.Email, &personnel.Website, &personnel.FileImage,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &personnel, nil
}

func (r *personnelRepository) CreatePersonnel(req models.PersonnelRequest) (*models.Personnels, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var academicID *int
	if req.AcademicPositionID != nil {
		academicID = req.AcademicPositionID
	} else if req.ThaiAcademicPosition != nil || req.EngAcademicPosition != nil {
		var id int
		err = tx.QueryRow(`SELECT academic_position_id FROM academic_position WHERE thai_academic_position = $1 AND eng_academic_position = $2`, req.ThaiAcademicPosition, req.EngAcademicPosition).Scan(&id)
		if err == sql.ErrNoRows {
			err = tx.QueryRow(`INSERT INTO academic_position (thai_academic_position, eng_academic_position) VALUES($1,$2) RETURNING academic_position_id`, req.ThaiAcademicPosition, req.EngAcademicPosition).Scan(&id)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		academicID = &id
	}

	var academicPositionID interface{}
	if academicID != nil {
		academicPositionID = *academicID
	} else {
		academicPositionID = nil
	}

	var newID int
	err = tx.QueryRow(`
		INSERT INTO personnels (
			type_personnel, department_position_id, academic_position_id,
			thai_name, eng_name, education, related_fields, email, website, file_image
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING personnel_id
	`,
		req.TypePersonnel, req.DepartmentPositionID, academicPositionID,
		req.ThaiName, req.EngName, req.Education, req.RelatedFields,
		req.Email, req.Website, req.FileImage,
	).Scan(&newID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.GetPersonnelByID(newID)
}

func (r *personnelRepository) UpdatePersonnel(id int, req models.PersonnelRequest) (*models.Personnels, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var academicID *int
	if req.AcademicPositionID != nil {
		academicID = req.AcademicPositionID
	} else if req.ThaiAcademicPosition != nil || req.EngAcademicPosition != nil {
		var aid int
		err = tx.QueryRow(`SELECT academic_position_id FROM academic_position WHERE thai_academic_position = $1 AND eng_academic_position = $2`, req.ThaiAcademicPosition, req.EngAcademicPosition).Scan(&aid)
		if err == sql.ErrNoRows {
			err = tx.QueryRow(`INSERT INTO academic_position (thai_academic_position, eng_academic_position) VALUES($1,$2) RETURNING academic_position_id`, req.ThaiAcademicPosition, req.EngAcademicPosition).Scan(&aid)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		academicID = &aid
	}

	var academicPositionID interface{}
	if academicID != nil {
		academicPositionID = *academicID
	} else {
		academicPositionID = nil
	}

	var updatedID int
	err = tx.QueryRow(`
		UPDATE personnels
		SET type_personnel=$1, department_position_id=$2, academic_position_id=$3,
			thai_name=$4, eng_name=$5, education=$6, related_fields=$7, email=$8, website=$9, file_image=$10
		WHERE personnel_id=$11
		RETURNING personnel_id
	`,
		req.TypePersonnel, req.DepartmentPositionID, academicPositionID,
		req.ThaiName, req.EngName, req.Education, req.RelatedFields,
		req.Email, req.Website, req.FileImage, id,
	).Scan(&updatedID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.GetPersonnelByID(updatedID)
}

func (r *personnelRepository) DeletePersonnel(id int) error {
	result, err := r.db.Exec("DELETE FROM personnels WHERE personnel_id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
