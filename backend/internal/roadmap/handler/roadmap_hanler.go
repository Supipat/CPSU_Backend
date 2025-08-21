package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/roadmap/models"
	"cpsu/internal/roadmap/service"

	"github.com/gin-gonic/gin"
)

type RoadmapHandler struct {
	cpsuService service.RoadmapService
}

func NewRoadmapHandler(cpsuService service.RoadmapService) *RoadmapHandler {
	return &RoadmapHandler{cpsuService: cpsuService}
}

func (h *RoadmapHandler) GetAllRoadmap(c *gin.Context) {
	var param models.RoadmapQueryParam
	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roadmaps, err := h.cpsuService.GetAllRoadmap(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roadmaps)
}

func (h *RoadmapHandler) GetRoadmapByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid roadmap ID"})
		return
	}

	roadmap, err := h.cpsuService.GetRoadmapByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "roadmap not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, roadmap)
}

func (h *RoadmapHandler) CreateRoadmap(c *gin.Context) {
	courseIDStr := c.PostForm("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil || courseID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	file, err := c.FormFile("roadmap_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roadmap_url is required"})
		return
	}

	created, err := h.cpsuService.CreateRoadmap(courseID, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *RoadmapHandler) DeleteRoadmap(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid roadmap ID"})
		return
	}

	err = h.cpsuService.DeleteRoadmap(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "roadmap not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "roadmap deleted successfully"})
}
