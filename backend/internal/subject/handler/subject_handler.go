package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/subject/models"
	"cpsu/internal/subject/service"

	"cpsu/internal/auth/repository"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	subjectService service.SubjectService
	auditRepo      *repository.AuditRepository
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
	var req models.SubjectsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdSubject, err := h.subjectService.CreateSubject(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "create", "subject",
		createdSubject.SubjectID,
		map[string]interface{}{
			"thai_subject": createdSubject.ThaiSubject,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

	c.JSON(http.StatusCreated, createdSubject)
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subject ID"})
		return
	}

	var req models.SubjectsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSubject, err := h.subjectService.UpdateSubject(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "subject not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "update", "subject",
		updatedSubject.SubjectID,
		map[string]interface{}{
			"thai_subject": updatedSubject.ThaiSubject,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

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

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "delete", "subject", strconv.Itoa(id),
		map[string]interface{}{
			"subject_id": id,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

	c.JSON(http.StatusOK, gin.H{"message": "subject deleted successfully"})
}
