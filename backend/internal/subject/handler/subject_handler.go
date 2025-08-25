package handler

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/subject/models"
	"cpsu/internal/subject/service"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	subjectService service.SubjectService
}

func NewSubjectHandler(subjectService service.SubjectService) *SubjectHandler {
	return &SubjectHandler{subjectService: subjectService}
}

func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {
	var param models.SubjectsQueryParam
	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subjects, err := h.subjectService.GetAllSubjects(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subjects)
}

func (h *SubjectHandler) GetSubjectByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subject ID"})
		return
	}

	subject, err := h.subjectService.GetSubjectByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "subject not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, subject)
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	subjectfile, err := c.FormFile("subjectfile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subjectfile required"})
		return
	}

	f, _ := subjectfile.Open()
	defer f.Close()
	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()

	var created []models.Subjects

	for i, row := range records {
		if i == 0 {
			continue
		}

		courseID, _ := strconv.Atoi(row[0])

		req := models.SubjectsRequest{
			CourseID: courseID, PlanType: row[1], Semester: row[2],
			ThaiSubject: row[3], EngSubject: row[4], Credits: row[5],
			CompulsorySubject: row[6], Condition: row[7], DescriptionThai: row[8],
			DescriptionEng: row[9], CLO: row[10],
		}

		createdSubject, err := h.subjectService.CreateSubject(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "row": i})
			return
		}
		created = append(created, *createdSubject)
	}

	c.JSON(http.StatusOK, created)
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subject ID"})
		return
	}

	courseIDStr := c.PostForm("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	req := models.SubjectsRequest{
		CourseID:          courseID,
		PlanType:          c.PostForm("plan_type"),
		Semester:          c.PostForm("semester"),
		ThaiSubject:       c.PostForm("thai_subject"),
		EngSubject:        c.PostForm("eng_subject"),
		Credits:           c.PostForm("credits"),
		CompulsorySubject: c.PostForm("compulsory_subject"),
		Condition:         c.PostForm("condition"),
		DescriptionThai:   c.PostForm("description_thai"),
		DescriptionEng:    c.PostForm("description_eng"),
		CLO:               c.PostForm("clo"),
	}

	updatedSubject, err := h.subjectService.UpdateSubject(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "subject ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, updatedSubject)
}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subject ID"})
		return
	}

	err = h.subjectService.DeleteSubject(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "subject not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subject deleted successfully"})
}
