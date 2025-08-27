package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/course/models"
	"cpsu/internal/course/service"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService service.CourseService
}

func NewCourseHandler(courseService service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	var param models.CoursesQueryParam
	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courses, err := h.courseService.GetAllCourses(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	course, err := h.courseService.GetCourseByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	yearStr := c.PostForm("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year"})
		return
	}

	req := models.CoursesRequest{
		Degree:        c.PostForm("degree"),
		Major:         c.PostForm("major"),
		Year:          year,
		ThaiCourse:    c.PostForm("thai_course"),
		EngCourse:     c.PostForm("eng_course"),
		ThaiDegree:    c.PostForm("thai_degree"),
		EngDegree:     c.PostForm("eng_degree"),
		AdmissionReq:  c.PostForm("admission_req"),
		GraduationReq: c.PostForm("graduation_req"),
		Philosophy:    c.PostForm("philosophy"),
		Objective:     c.PostForm("objective"),
		Tuition:       c.PostForm("tuition"),
		Credits:       c.PostForm("credits"),
		CareerPaths:   c.PostForm("career_paths"),
		PLO:           c.PostForm("plo"),
		DetailURL:     c.PostForm("detail_url"),
	}

	createdCourse, err := h.courseService.CreateCourse(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCourse)
}

/*func (h *CourseHandler) CreateCourse(c *gin.Context) {
	coursefile, err := c.FormFile("coursefile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "coursefile required"})
		return
	}

	f, _ := coursefile.Open()
	defer f.Close()
	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()

	var created []models.Courses

	for i, row := range records {
		if i == 0 {
			continue
		}
		year, _ := strconv.Atoi(row[2])
		req := models.CoursesRequest{
			Degree: row[0], Major: row[1], Year: year,
			ThaiCourse: row[3], EngCourse: row[4],
			ThaiDegree: row[5], EngDegree: row[6],
			AdmissionReq: row[7], GraduationReq: row[8],
			Philosophy: row[9], Objective: row[10],
			Tuition: row[11], Credits: row[12],
			CareerPaths: row[13], PLO: row[14], DetailURL: row[15],
		}
		createdCourse, err := h.courseService.CreateCourse(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "row": i})
			return
		}
		created = append(created, *createdCourse)
	}

	c.JSON(http.StatusOK, created)
}*/

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	yearStr := c.PostForm("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year"})
		return
	}

	req := models.CoursesRequest{
		Degree:        c.PostForm("degree"),
		Major:         c.PostForm("major"),
		Year:          year,
		ThaiCourse:    c.PostForm("thai_course"),
		EngCourse:     c.PostForm("eng_course"),
		ThaiDegree:    c.PostForm("thai_degree"),
		EngDegree:     c.PostForm("eng_degree"),
		AdmissionReq:  c.PostForm("admission_req"),
		GraduationReq: c.PostForm("graduation_req"),
		Philosophy:    c.PostForm("philosophy"),
		Objective:     c.PostForm("objective"),
		Tuition:       c.PostForm("tuition"),
		Credits:       c.PostForm("credits"),
		CareerPaths:   c.PostForm("career_paths"),
		PLO:           c.PostForm("plo"),
		DetailURL:     c.PostForm("detail_url"),
	}

	updatedCourse, err := h.courseService.UpdateCourse(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, updatedCourse)
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	err = h.courseService.DeleteCourse(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "course deleted successfully"})
}
