package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"cpsu/internal/course_structure/models"
)

type CourseStructureRepository interface {
	GetAllCourseStructure(param models.CourseStructureQueryParam) ([]models.CourseStructure, error)
	GetCourseStructureByID(id int) (*models.CourseStructure, error)
	CreateCourseStructure(req *models.CourseStructureRequest) (*models.CourseStructure, error)
	DeleteCourseStructure(id int) error
}

type courseStructureRepository struct {
	db *sql.DB
}

func NewCourseStructureRepository(db *sql.DB) CourseStructureRepository {
	return &courseStructureRepository{db: db}
}

func (r *courseStructureRepository) GetAllCourseStructure(param models.CourseStructureQueryParam) ([]models.CourseStructure, error) {
	query := `
		SELECT cs.course_structure_id, c.course_id, c.thai_course, cs.course_structure_url
		FROM course_structure cs
		LEFT JOIN courses c ON cs.course_id = c.course_id
	`

	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if param.CourseID > 0 {
		conditions = append(conditions, "c.course_id = $"+strconv.Itoa(argIndex))
		args = append(args, param.CourseID)
		argIndex++
	}

	if len(param.Search) > 0 {
		conditions = append(conditions, "c.thai_course ILIKE '%' || $"+strconv.Itoa(argIndex)+" || '%'")
		args = append(args, param.Search)
		argIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	sort := "cs.course_structure_id"
	if param.Sort != "" {
		sort = "cs." + param.Sort
	}

	order := "DESC"
	if strings.ToUpper(param.Order) == "ASC" {
		order = "ASC"
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

	var courseStructures []models.CourseStructure
	for rows.Next() {
		var cs models.CourseStructure
		err := rows.Scan(&cs.CourseStructureID, &cs.CourseID, &cs.ThaiCourse, &cs.CourseStructureURL)
		if err != nil {
			return nil, err
		}
		courseStructures = append(courseStructures, cs)
	}

	return courseStructures, nil
}

func (r *courseStructureRepository) GetCourseStructureByID(id int) (*models.CourseStructure, error) {
	query := `
		SELECT cs.course_structure_id, c.course_id, c.thai_course, cs.course_structure_url
		FROM course_structure cs
		LEFT JOIN courses c ON cs.course_id = c.course_id
		WHERE cs.course_structure_id = $1
	`
	row := r.db.QueryRow(query, id)

	var cs models.CourseStructure
	err := row.Scan(&cs.CourseStructureID, &cs.CourseID, &cs.ThaiCourse, &cs.CourseStructureURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &cs, nil
}

func (r *courseStructureRepository) CreateCourseStructure(req *models.CourseStructureRequest) (*models.CourseStructure, error) {
	query := `
		INSERT INTO course_structure (course_id, course_structure_url)
		VALUES ($1, $2)
		RETURNING course_structure_id
	`

	var cs models.CourseStructure
	err := r.db.QueryRow(query, req.CourseID, req.CourseStructureURL).Scan(&cs.CourseStructureID)
	if err != nil {
		return nil, err
	}

	cs.CourseID = req.CourseID
	cs.CourseStructureURL = req.CourseStructureURL

	row := r.db.QueryRow("SELECT thai_course FROM courses WHERE course_id = $1", req.CourseID)
	_ = row.Scan(&cs.ThaiCourse)

	return &cs, nil
}

func (r *courseStructureRepository) DeleteCourseStructure(id int) error {
	result, err := r.db.Exec("DELETE FROM course_structure WHERE course_structure_id = $1", id)
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
