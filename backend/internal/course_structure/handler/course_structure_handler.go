package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/course_structure/models"
	"cpsu/internal/course_structure/service"

	"github.com/gin-gonic/gin"
)

type CourseStructureHandler struct {
	courseStructureService service.CourseStructureService
}

func NewCourseStructureHandler(courseStructureService service.CourseStructureService) *CourseStructureHandler {
	return &CourseStructureHandler{courseStructureService: courseStructureService}
}

func (h *CourseStructureHandler) GetAllCourseStructure(c *gin.Context) {
	var param models.CourseStructureQueryParam
	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courseStructures, err := h.courseStructureService.GetAllCourseStructure(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courseStructures)
}

func (h *CourseStructureHandler) GetCourseStructureByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course_structure ID"})
		return
	}

	cs, err := h.courseStructureService.GetCourseStructureByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course_structure not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, cs)
}

func (h *CourseStructureHandler) CreateCourseStructure(c *gin.Context) {
	courseIDStr := c.PostForm("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil || courseID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	file, err := c.FormFile("course_structure_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course_structure_url is required"})
		return
	}

	created, err := h.courseStructureService.CreateCourseStructure(courseID, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *CourseStructureHandler) DeleteCourseStructure(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course_structure ID"})
		return
	}

	err = h.courseStructureService.DeleteCourseStructure(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course_structure not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "course_structure deleted successfully"})
}
