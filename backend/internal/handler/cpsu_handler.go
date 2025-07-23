package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/service"

	"github.com/gin-gonic/gin"
)

type CPSUHandler struct {
	cpsuService service.CPSUService
}

func NewCPSUHandler(cpsuService service.CPSUService) *CPSUHandler {
	return &CPSUHandler{cpsuService: cpsuService}
}

func (h *CPSUHandler) GetAllNews(c *gin.Context) {
	var param service.NewsQueryParam
	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameter"})
		return
	}

	newsList, err := h.cpsuService.GetAllNews(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get news"})
		return
	}
	c.JSON(http.StatusOK, newsList)
}

func (h *CPSUHandler) GetNewsDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid news ID"})
		return
	}

	news, err := h.cpsuService.GetNewsDetail(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "news ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, news)
}

type NewsRequest struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	NewsType  string   `json:"news_type"`
	DetailURL string   `json:"detail_url"`
	Images    []string `json:"images"`
}

func (h *CPSUHandler) CreateNews(c *gin.Context) {
	var req NewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	news, err := h.cpsuService.CreateNews(req.Title, req.Content, req.NewsType, req.DetailURL, req.Images)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, news)
}

func (h *CPSUHandler) UpdateNews(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid news ID"})
		return
	}

	var req NewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	updated, err := h.cpsuService.UpdateNews(id, req.Title, req.Content, req.NewsType, req.DetailURL, req.Images)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "news ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *CPSUHandler) DeleteNews(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid news ID"})
		return
	}

	err = h.cpsuService.DeleteNews(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "news ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}
