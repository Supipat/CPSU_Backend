package repository

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"cpsu/internal/personnel/models"
)

type PersonnelRepository interface {
	GetAllPersonnels(param models.PersonnelQueryParam) ([]models.Personnels, error)
	GetPersonnelByID(id int) (*models.Personnels, error)
	CreatePersonnel(req models.PersonnelRequest) (*models.Personnels, error)
	UpdatePersonnel(id int, req models.PersonnelRequest) (*models.Personnels, error)
	DeletePersonnel(id int) error
	GetScopusIDByPersonnelID(id int) (*string, error)
	SaveResearch(personnelID int, researches []models.Research) (err error)
	GetAllResearch(param models.ResearchQueryParam) ([]models.Research, error)
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
			p.eng_name, p.education, p.related_fields, p.email, p.website, p.file_image, p.scopus_id
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
		var scopus sql.NullString
		err := rows.Scan(
			&personnel.PersonnelID, &personnel.TypePersonnel, &personnel.DepartmentPositionID, &personnel.DepartmentPositionName,
			&personnel.AcademicPositionID, &personnel.ThaiAcademicPosition, &personnel.EngAcademicPosition,
			&personnel.ThaiName, &personnel.EngName, &personnel.Education, &personnel.RelatedFields,
			&personnel.Email, &personnel.Website, &personnel.FileImage, &scopus,
		)
		if err != nil {
			return nil, err
		}
		if scopus.Valid {
			personnel.ScopusID = &scopus.String
		} else {
			personnel.ScopusID = nil
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
			p.eng_name, p.education, p.related_fields, p.email, p.website, p.file_image, p.scopus_id
		FROM personnels p
		LEFT JOIN department_position d ON p.department_position_id = d.department_position_id
		LEFT JOIN academic_position a ON p.academic_position_id = a.academic_position_id
		WHERE p.personnel_id = $1
	`

	row := r.db.QueryRow(query, id)

	var personnel models.Personnels
	var scopus sql.NullString
	err := row.Scan(
		&personnel.PersonnelID, &personnel.TypePersonnel, &personnel.DepartmentPositionID, &personnel.DepartmentPositionName,
		&personnel.AcademicPositionID, &personnel.ThaiAcademicPosition, &personnel.EngAcademicPosition,
		&personnel.ThaiName, &personnel.EngName, &personnel.Education, &personnel.RelatedFields,
		&personnel.Email, &personnel.Website, &personnel.FileImage, &scopus,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	if scopus.Valid {
		personnel.ScopusID = &scopus.String
	} else {
		personnel.ScopusID = nil
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
			type_personnel, department_position_id, academic_position_id,thai_name, 
			eng_name, education, related_fields, email, website, file_image, scopus_id
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING personnel_id
	`,
		req.TypePersonnel, req.DepartmentPositionID, academicPositionID,
		req.ThaiName, req.EngName, req.Education, req.RelatedFields,
		req.Email, req.Website, req.FileImage, req.ScopusID,
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
		SET type_personnel=$1, department_position_id=$2, academic_position_id=$3,thai_name=$4, eng_name=$5, 
		education=$6, related_fields=$7, email=$8, website=$9, file_image=$10, scopus_id=$11
		WHERE personnel_id=$12
		RETURNING personnel_id
	`,
		req.TypePersonnel, req.DepartmentPositionID, academicPositionID,
		req.ThaiName, req.EngName, req.Education, req.RelatedFields,
		req.Email, req.Website, req.FileImage, req.ScopusID, id,
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

func (r *personnelRepository) GetScopusIDByPersonnelID(id int) (*string, error) {
	query := `SELECT scopus_id FROM personnels WHERE personnel_id = $1`
	var scopus sql.NullString
	err := r.db.QueryRow(query, id).Scan(&scopus)
	if err != nil {
		return nil, err
	}
	if scopus.Valid {
		return &scopus.String, nil
	}
	return nil, nil
}

func (r *personnelRepository) SaveResearch(personnelID int, researches []models.Research) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	val := func(s *string) interface{} {
		if s == nil {
			return nil
		}
		return *s
	}

	for _, rc := range researches {
		if rc.DOI != nil && *rc.DOI != "" {
			res, e := tx.Exec(
				`UPDATE research SET title=$1, journal=$2, year=$3, volume=$4, issue=$5, pages=$6, cited=$7
                 WHERE personnel_id=$8 AND doi=$9`,
				rc.Title, rc.Journal, rc.Year, val(rc.Volume), val(rc.Issue), val(rc.Pages), rc.Cited, personnelID, *rc.DOI,
			)
			if e != nil {
				return e
			}
			if rows, _ := res.RowsAffected(); rows > 0 {
				continue
			}
		}

		res, e := tx.Exec(
			`UPDATE research SET journal=$1, volume=$2, issue=$3, pages=$4, doi=$5, cited=$6
             WHERE personnel_id=$7 AND title=$8 AND year=$9`,
			rc.Journal, val(rc.Volume), val(rc.Issue), val(rc.Pages), val(rc.DOI), rc.Cited, personnelID, rc.Title, rc.Year,
		)
		if e != nil {
			return e
		}
		if rows, _ := res.RowsAffected(); rows > 0 {
			continue
		}

		createdAt := rc.CreatedAt
		if createdAt.IsZero() {
			createdAt = time.Now()
		}
		if _, e = tx.Exec(
			`INSERT INTO research (personnel_id, title, journal, year, volume, issue, pages, doi, cited, created_at)
             VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
			personnelID, rc.Title, rc.Journal, rc.Year, val(rc.Volume), val(rc.Issue), val(rc.Pages), val(rc.DOI), rc.Cited, createdAt,
		); e != nil {
			return e
		}
	}
	return nil
}

func (r *personnelRepository) GetAllResearch(param models.ResearchQueryParam) ([]models.Research, error) {
	query := `
		SELECT r.research_id, p.personnel_id, r.title, r.journal, r.year, r.volume, r.issue, r.pages, r.doi, r.cited, r.created_at
        FROM research r
		LEFT JOIN personnels p ON r.personnel_id = p.personnel_id
	`
	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if param.PersonnelID > 0 {
		conditions = append(conditions, "p.personnel_id = $"+strconv.Itoa(argIndex))
		args = append(args, param.PersonnelID)
		argIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	sort := "research_id"
	if param.Sort != "" {
		sort = param.Sort
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

	var researches []models.Research
	for rows.Next() {
		var r models.Research
		var vol, iss, pages, doi sql.NullString

		err := rows.Scan(
			&r.ResearchID, &r.PersonnelID, &r.Title, &r.Journal, &r.Year,
			&vol, &iss, &pages, &doi, &r.Cited, &r.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if vol.Valid {
			r.Volume = &vol.String
		}
		if iss.Valid {
			r.Issue = &iss.String
		}
		if pages.Valid {
			r.Pages = &pages.String
		}
		if doi.Valid {
			r.DOI = &doi.String
		}

		researches = append(researches, r)
	}

	return researches, nil
}
