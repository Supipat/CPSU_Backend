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
			c.course_id, d.degree_id, d.degree, m.major_id, m.major, c.year, c.thai_course, 
			c.eng_course, dn.degree_name_id, dn.thai_degree, dn.eng_degree, c.admission_req, 
			c.graduation_req, c.philosophy, c.objective, c.tuition, c.credits, 
			cp.career_paths_id, cp.career_paths, p.plo_id, p.plo, c.detail_url
		FROM courses c
		LEFT JOIN degree d ON c.degree_id = d.degree_id
		LEFT JOIN majors m ON c.major_id = m.major_id
		LEFT JOIN degree_name dn ON c.degree_name_id = dn.degree_name_id
		LEFT JOIN career_paths cp ON c.career_paths_id = cp.career_paths_id
		LEFT JOIN plo p ON c.plo_id = p.plo_id
	`

	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if param.DegreeID > 0 {
		conditions = append(conditions, "d.degree_id = $"+strconv.Itoa(argIndex))
		args = append(args, param.DegreeID)
		argIndex++
	}

	if param.MajorID > 0 {
		conditions = append(conditions, "m.major_id = $"+strconv.Itoa(argIndex))
		args = append(args, param.MajorID)
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
			&course.CourseID, &course.DegreeID, &course.Degree, &course.MajorID, &course.Major,
			&course.Year, &course.ThaiCourse, &course.EngCourse, &course.DegreeNameID,
			&course.ThaiDegree, &course.EngDegree, &course.AdmissionReq, &course.GraduationReq,
			&course.Philosophy, &course.Objective, &course.Tuition, &course.Credits,
			&course.CareerPathsID, &course.CareerPaths, &course.PloID,
			&course.PLO, &course.DetailURL,
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
			c.course_id, d.degree_id, d.degree, m.major_id, m.major, c.year, c.thai_course, 
			c.eng_course, dn.degree_name_id, dn.thai_degree, dn.eng_degree, c.admission_req, 
			c.graduation_req, c.philosophy, c.objective, c.tuition, c.credits, 
			cp.career_paths_id, cp.career_paths, p.plo_id, p.plo, c.detail_url
		FROM courses c
		LEFT JOIN degree d ON c.degree_id = d.degree_id
		LEFT JOIN majors m ON c.major_id = m.major_id
		LEFT JOIN degree_name dn ON c.degree_name_id = dn.degree_name_id
		LEFT JOIN career_paths cp ON c.career_paths_id = cp.career_paths_id
		LEFT JOIN plo p ON c.plo_id = p.plo_id
		WHERE c.course_id = $1
	`

	row := r.db.QueryRow(query, id)

	var course models.Courses
	err := row.Scan(
		&course.CourseID, &course.DegreeID, &course.Degree, &course.MajorID, &course.Major,
		&course.Year, &course.ThaiCourse, &course.EngCourse, &course.DegreeNameID,
		&course.ThaiDegree, &course.EngDegree, &course.AdmissionReq, &course.GraduationReq,
		&course.Philosophy, &course.Objective, &course.Tuition, &course.Credits,
		&course.CareerPathsID, &course.CareerPaths, &course.PloID,
		&course.PLO, &course.DetailURL,
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
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var majorID int
	err = tx.QueryRow(`SELECT major_id FROM majors WHERE major = $1`, req.Major).Scan(&majorID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO majors (major) VALUES($1) RETURNING major_id`, req.Major).Scan(&majorID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	var degreeNameID int
	err = tx.QueryRow(`SELECT degree_name_id FROM degree_name WHERE thai_degree = $1`, req.ThaiDegree).Scan(&degreeNameID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO degree_name (thai_degree,eng_degree) VALUES($1,$2) RETURNING degree_name_id`, req.ThaiDegree, req.EngDegree).Scan(&degreeNameID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	var careerPathsID int
	err = tx.QueryRow(`SELECT career_paths_id FROM career_paths WHERE career_paths = $1`, req.CareerPaths).Scan(&careerPathsID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO career_paths (career_paths) VALUES($1) RETURNING career_paths_id`, req.CareerPaths).Scan(&careerPathsID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	var ploID int
	err = tx.QueryRow(`SELECT plo_id FROM plo WHERE plo = $1`, req.PLO).Scan(&ploID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO plo (plo) VALUES($1) RETURNING plo_id`, req.PLO).Scan(&ploID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	var courseID int
	err = tx.QueryRow(`
        INSERT INTO courses (
            degree_id, major_id, year, thai_course, eng_course, degree_name_id,
            admission_req, graduation_req, philosophy, objective, tuition, credits,
            career_paths_id, plo_id, detail_url
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
        RETURNING course_id
    `,
		req.DegreeID, majorID, req.Year, req.ThaiCourse, req.EngCourse,
		degreeNameID, req.AdmissionReq, req.GraduationReq, req.Philosophy,
		req.Objective, req.Tuition, req.Credits, careerPathsID, ploID, req.DetailURL,
	).Scan(&courseID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.GetCourseByID(courseID)
}

func (r *courseRepository) UpdateCourse(id int, req models.CoursesRequest) (*models.Courses, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var majorID int
	err = tx.QueryRow(`SELECT major_id FROM majors WHERE major = $1`, req.Major).Scan(&majorID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO majors (major) VALUES($1) RETURNING major_id`, req.Major).Scan(&majorID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	var degreeNameID int
	err = tx.QueryRow(`SELECT degree_name_id FROM degree_name WHERE thai_degree = $1`, req.ThaiDegree).Scan(&degreeNameID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO degree_name (thai_degree,eng_degree) VALUES($1,$2) RETURNING degree_name_id`, req.ThaiDegree, req.EngDegree).Scan(&degreeNameID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	var careerPathsID int
	err = tx.QueryRow(`SELECT career_paths_id FROM career_paths WHERE career_paths = $1`, req.CareerPaths).Scan(&careerPathsID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO career_paths (career_paths) VALUES($1) RETURNING career_paths_id`, req.CareerPaths).Scan(&careerPathsID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	var ploID int
	err = tx.QueryRow(`SELECT plo_id FROM plo WHERE plo = $1`, req.PLO).Scan(&ploID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO plo (plo) VALUES($1) RETURNING plo_id`, req.PLO).Scan(&ploID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		UPDATE courses
		SET degree_id=$1, major_id=$2, year=$3, thai_course=$4, eng_course=$5, degree_name_id=$6,
		    admission_req=$7, graduation_req=$8, philosophy=$9, objective=$10, tuition=$11,
		    credits=$12, career_paths_id=$13, plo_id=$14, detail_url=$15
		WHERE course_id=$16
	`,
		req.DegreeID, majorID, req.Year, req.ThaiCourse, req.EngCourse, degreeNameID,
		req.AdmissionReq, req.GraduationReq, req.Philosophy, req.Objective, req.Tuition,
		req.Credits, careerPathsID, ploID, req.DetailURL, id,
	)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.GetCourseByID(id)
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
