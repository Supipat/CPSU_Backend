package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/document/models"
	"cpsu/internal/document/service"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	documentService service.DocumentService
}

func NewDocumentHandler(documentService service.DocumentService) *DocumentHandler {
	return &DocumentHandler{documentService: documentService}
}

func (h *DocumentHandler) GetAllDocument(c *gin.Context) {

	var param models.DocumentQueryParam

	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameter"})
		return
	}

	documents, err := h.documentService.GetAllDocument(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get document"})
		return
	}

	c.JSON(http.StatusOK, documents)
}

func (h *DocumentHandler) GetDocumentByID(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document ID"})
		return
	}

	document, err := h.documentService.GetDocumentByID(id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "document ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, document)
}

func (h *DocumentHandler) CreateDocument(c *gin.Context) {

	title := c.PostForm("title")
	typeIDStr := c.PostForm("type_id")
	typeName := c.PostForm("type_name")

	var description *string
	descriptionStr := c.PostForm("description")

	if descriptionStr != "" {
		description = &descriptionStr
	}

	var typeID int

	if typeIDStr != "" {
		parsedID, err := strconv.Atoi(typeIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type ID"})
			return
		}

		typeID = parsedID
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	userID := c.GetInt("user_id")
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	created, err := h.documentService.CreateDocument(
		typeID, typeName, title, description,
		file, userID, ip, userAgent,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *DocumentHandler) UpdateDocument(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document ID"})
		return
	}

	title := c.PostForm("title")
	typeIDStr := c.PostForm("type_id")
	typeName := c.PostForm("type_name")

	var description *string
	descriptionStr := c.PostForm("description")

	if descriptionStr != "" {
		description = &descriptionStr
	}

	var typeID int

	if typeIDStr != "" {
		parsedID, err := strconv.Atoi(typeIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type ID"})
			return
		}

		typeID = parsedID
	}

	file, _ := c.FormFile("file")

	userID := c.GetInt("user_id")
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	updated, err := h.documentService.UpdateDocument(
		id, typeID, typeName, title, description,
		file, userID, ip, userAgent,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "document ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *DocumentHandler) DeleteDocument(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document ID"})
		return
	}

	userID := c.GetInt("user_id")
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	err = h.documentService.DeleteDocument(id, userID, ip, userAgent)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "document ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
