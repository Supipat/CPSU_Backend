package handler

import (
	"database/sql"
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

	subject, err := h.subjectService.GetAllSubjects(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subject)
}

func (h *SubjectHandler) GetSubjectByID(c *gin.Context) {
	subjectID := c.Param("id")
	if subjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject ID required"})
		return
	}

	subject, err := h.subjectService.GetSubjectByID(subjectID)
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
	courseIDStr := c.PostForm("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course_id"})
		return
	}

	Ptr := func(s string) *string {
		if s == "" {
			return nil
		}
		return &s
	}

	req := models.SubjectsRequest{
		SubjectID:         c.PostForm("subject_id"),
		CourseID:          courseID,
		PlanType:          c.PostForm("plan_type"),
		Semester:          c.PostForm("semester"),
		ThaiSubject:       c.PostForm("thai_subject"),
		EngSubject:        Ptr(c.PostForm("eng_subject")),
		Credits:           c.PostForm("credits"),
		CompulsorySubject: Ptr(c.PostForm("compulsory_subject")),
		Condition:         Ptr(c.PostForm("condition")),
		DescriptionThai:   Ptr(c.PostForm("description_thai")),
		DescriptionEng:    Ptr(c.PostForm("description_eng")),
		CLO:               Ptr(c.PostForm("clo")),
	}

	createdSubject, err := h.subjectService.CreateSubject(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSubject)
}

/*func (h *SubjectHandler) CreateSubject(c *gin.Context) {
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

		courseID, _ := strconv.Atoi(row[1])

		Ptr := func(s string) *string {
			if s == "" {
				return nil
			}
			return &s
		}

		req := models.SubjectsRequest{
			SubjectID: row[0], CourseID: courseID, PlanType: row[2], Semester: row[3],
			ThaiSubject: row[4], EngSubject: Ptr(row[5]), Credits: row[6], CompulsorySubject: Ptr(row[7]),
			Condition: Ptr(row[8]), DescriptionThai: Ptr(row[9]), DescriptionEng: Ptr(row[10]), CLO: Ptr(row[11]),
		}

		createdSubject, err := h.subjectService.CreateSubject(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "row": i})
			return
		}
		created = append(created, *createdSubject)
	}

	c.JSON(http.StatusOK, created)
}*/

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	subjectID := c.Param("id")
	if subjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject ID required"})
		return
	}

	courseIDStr := c.PostForm("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	Ptr := func(s string) *string {
		if s == "" {
			return nil
		}
		return &s
	}

	req := models.SubjectsRequest{
		CourseID:          courseID,
		PlanType:          c.PostForm("plan_type"),
		Semester:          c.PostForm("semester"),
		ThaiSubject:       c.PostForm("thai_subject"),
		EngSubject:        Ptr(c.PostForm("eng_subject")),
		Credits:           c.PostForm("credits"),
		CompulsorySubject: Ptr(c.PostForm("compulsory_subject")),
		Condition:         Ptr(c.PostForm("condition")),
		DescriptionThai:   Ptr(c.PostForm("description_thai")),
		DescriptionEng:    Ptr(c.PostForm("description_eng")),
		CLO:               Ptr(c.PostForm("clo")),
	}

	updatedSubject, err := h.subjectService.UpdateSubject(subjectID, req)
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
	subjectID := c.Param("id")
	if subjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject ID required"})
		return
	}

	err := h.subjectService.DeleteSubject(subjectID)
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
