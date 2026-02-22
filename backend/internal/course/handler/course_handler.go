package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"cpsu/internal/course/models"
	"cpsu/internal/course/service"

	"cpsu/internal/auth/repository"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService service.CourseService
	auditRepo     *repository.AuditRepository
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
	id := c.Param("id")

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
	var req models.CoursesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCourse, err := h.courseService.CreateCourse(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "create", "course",
		createdCourse.CourseID,
		map[string]interface{}{
			"thai_course": createdCourse.ThaiCourse,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

	c.JSON(http.StatusCreated, createdCourse)
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	var req models.CoursesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCourse, err := h.courseService.UpdateCourse(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "update", "course",
		updatedCourse.CourseID,
		map[string]interface{}{
			"thai_course": updatedCourse.ThaiCourse,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

	c.JSON(http.StatusOK, updatedCourse)
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	err := h.courseService.DeleteCourse(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "delete", "news", id,
		map[string]interface{}{
			"course_id": id,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

	c.JSON(http.StatusOK, gin.H{"message": "course deleted successfully"})
}
