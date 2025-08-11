package handler

import (
	"database/sql"
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"

	"cpsu/internal/news/models"
	"cpsu/internal/news/service"

	"github.com/gin-gonic/gin"
)

type CPSUHandler struct {
	cpsuService service.CPSUService
}

func NewCPSUHandler(cpsuService service.CPSUService) *CPSUHandler {
	return &CPSUHandler{cpsuService: cpsuService}
}

func (h *CPSUHandler) GetAllNews(c *gin.Context) {
	var param models.NewsQueryParam
	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameter"})
		return
	}

	newsList, err := h.cpsuService.GetAllNews(param)
	if err != nil {
		if err.Error() == "news type not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get news"})
		}
		return
	}
	c.JSON(http.StatusOK, newsList)
}

func (h *CPSUHandler) GetNewsByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid news ID"})
		return
	}

	news, err := h.cpsuService.GetNewsByID(id)
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

func (h *CPSUHandler) CreateNews(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	typeIDStr := c.PostForm("type_id")
	detailURL := c.PostForm("detail_url")

	typeID, err := strconv.Atoi(typeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type ID"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse form"})
		return
	}

	var fileImages []*multipart.FileHeader
	if form != nil {
		fileImages = form.File["images"]
	}

	created, err := h.cpsuService.CreateNews(title, content, typeID, "", detailURL, fileImages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *CPSUHandler) UpdateNews(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid news ID"})
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")
	typeIDStr := c.PostForm("type_id")
	detailURL := c.PostForm("detail_url")

	typeID, err := strconv.Atoi(typeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type ID"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse form"})
		return
	}

	var fileImages []*multipart.FileHeader
	if form != nil {
		fileImages = form.File["images"]
	}

	updated, err := h.cpsuService.UpdateNews(id, title, content, typeID, "", detailURL, fileImages)
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
