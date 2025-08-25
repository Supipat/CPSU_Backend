package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"cpsu/internal/course/models"
)

type CourseRepository interface {
	GetAllCourses(param models.CoursesQueryParam) ([]models.Courses, error)
	GetCourseByID(id int) (*models.Courses, error)
	CreateCourse(req models.CoursesRequest) (*models.Courses, error)
	UpdateCourse(id int, req models.CoursesRequest) (*models.Courses, error)
	DeleteCourse(id int) error
}

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) GetAllCourses(param models.CoursesQueryParam) ([]models.Courses, error) {
	query := `
		SELECT 
			c.course_id, c.degree, c.major, c.year,c.thai_course, 
			c.eng_course,c.thai_degree, c.eng_degree,c.admission_req, 
			c.graduation_req,c.philosophy, c.objective,c.tuition, 
			c.credits,c.career_paths, c.plo, c.detail_url
		FROM courses c
	`

	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if param.Degree != "" {
		conditions = append(conditions, "c.degree = $"+strconv.Itoa(argIndex))
		args = append(args, param.Degree)
		argIndex++
	}

	if param.Major != "" {
		conditions = append(conditions, "c.major = $"+strconv.Itoa(argIndex))
		args = append(args, param.Major)
		argIndex++
	}

	if param.Year > 0 {
		conditions = append(conditions, "c.year = $"+strconv.Itoa(argIndex))
		args = append(args, param.Year)
		argIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	sort := "c.year"
	if param.Sort != "" {
		sort = "c." + param.Sort
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

	var courses []models.Courses
	for rows.Next() {
		var course models.Courses
		err := rows.Scan(
			&course.CourseID, &course.Degree, &course.Major, &course.Year, &course.ThaiCourse,
			&course.EngCourse, &course.ThaiDegree, &course.EngDegree, &course.AdmissionReq,
			&course.GraduationReq, &course.Philosophy, &course.Objective, &course.Tuition,
			&course.Credits, &course.CareerPaths, &course.PLO, &course.DetailURL,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (r *courseRepository) GetCourseByID(id int) (*models.Courses, error) {
	query := `
		SELECT 
			c.course_id, c.degree, c.major, c.year,c.thai_course, 
			c.eng_course,c.thai_degree, c.eng_degree,c.admission_req, 
			c.graduation_req,c.philosophy, c.objective,c.tuition, 
			c.credits,c.career_paths, c.plo, c.detail_url
		FROM courses c
		WHERE c.course_id = $1
	`

	row := r.db.QueryRow(query, id)

	var course models.Courses
	err := row.Scan(
		&course.CourseID, &course.Degree, &course.Major, &course.Year, &course.ThaiCourse,
		&course.EngCourse, &course.ThaiDegree, &course.EngDegree, &course.AdmissionReq,
		&course.GraduationReq, &course.Philosophy, &course.Objective, &course.Tuition,
		&course.Credits, &course.CareerPaths, &course.PLO, &course.DetailURL,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &course, nil
}

func (r *courseRepository) CreateCourse(req models.CoursesRequest) (*models.Courses, error) {
	query := `
		INSERT INTO courses (
			degree, major, year, thai_course, eng_course, thai_degree, eng_degree,
			admission_req, graduation_req, philosophy, objective, tuition,
			credits, career_paths, plo, detail_url
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
		RETURNING course_id
	`

	var Course models.Courses
	err := r.db.QueryRow(
		query,
		req.Degree, req.Major, req.Year, req.ThaiCourse, req.EngCourse,
		req.ThaiDegree, req.EngDegree, req.AdmissionReq, req.GraduationReq,
		req.Philosophy, req.Objective, req.Tuition, req.Credits,
		req.CareerPaths, req.PLO, req.DetailURL,
	).Scan(&Course.CourseID)
	if err != nil {
		return nil, err
	}

	Course.Degree = req.Degree
	Course.Major = req.Major
	Course.Year = req.Year
	Course.ThaiCourse = req.ThaiCourse
	Course.EngCourse = req.EngCourse
	Course.ThaiDegree = req.ThaiDegree
	Course.EngDegree = req.EngDegree
	Course.AdmissionReq = req.AdmissionReq
	Course.GraduationReq = req.GraduationReq
	Course.Philosophy = req.Philosophy
	Course.Objective = req.Objective
	Course.Tuition = req.Tuition
	Course.Credits = req.Credits
	Course.CareerPaths = req.CareerPaths
	Course.PLO = req.PLO
	Course.DetailURL = req.DetailURL

	return &Course, nil
}

func (r *courseRepository) UpdateCourse(id int, req models.CoursesRequest) (*models.Courses, error) {
	query := `
		UPDATE courses
		SET degree=$1, major=$2, year=$3, thai_course=$4, eng_course=$5,
			thai_degree=$6, eng_degree=$7, admission_req=$8, graduation_req=$9,
			philosophy=$10, objective=$11, tuition=$12, credits=$13,
			career_paths=$14, plo=$15, detail_url=$16
		WHERE course_id=$17
		RETURNING course_id
	`

	var Course models.Courses
	err := r.db.QueryRow(
		query,
		req.Degree, req.Major, req.Year, req.ThaiCourse, req.EngCourse,
		req.ThaiDegree, req.EngDegree, req.AdmissionReq, req.GraduationReq,
		req.Philosophy, req.Objective, req.Tuition, req.Credits,
		req.CareerPaths, req.PLO, req.DetailURL,
		id,
	).Scan(&Course.CourseID)
	if err != nil {
		return nil, err
	}

	Course.Degree = req.Degree
	Course.Major = req.Major
	Course.Year = req.Year
	Course.ThaiCourse = req.ThaiCourse
	Course.EngCourse = req.EngCourse
	Course.ThaiDegree = req.ThaiDegree
	Course.EngDegree = req.EngDegree
	Course.AdmissionReq = req.AdmissionReq
	Course.GraduationReq = req.GraduationReq
	Course.Philosophy = req.Philosophy
	Course.Objective = req.Objective
	Course.Tuition = req.Tuition
	Course.Credits = req.Credits
	Course.CareerPaths = req.CareerPaths
	Course.PLO = req.PLO
	Course.DetailURL = req.DetailURL

	return &Course, nil
}

func (r *courseRepository) DeleteCourse(id int) error {
	result, err := r.db.Exec("DELETE FROM courses WHERE course_id = $1", id)
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
