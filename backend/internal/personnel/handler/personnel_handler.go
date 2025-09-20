package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/personnel/models"
	"cpsu/internal/personnel/service"

	"github.com/gin-gonic/gin"
)

type PersonnelHandler struct {
	personnelService service.PersonnelService
}

func NewPersonnelHandler(personnelService service.PersonnelService) *PersonnelHandler {
	return &PersonnelHandler{personnelService: personnelService}
}

func (h *PersonnelHandler) GetAllPersonnels(c *gin.Context) {
	var param models.PersonnelQueryParam
	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameter"})
		return
	}

	personnels, err := h.personnelService.GetAllPersonnels(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get personnels"})
		return
	}

	c.JSON(http.StatusOK, personnels)
}

func (h *PersonnelHandler) GetPersonnelByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid personnel ID"})
		return
	}

	personnel, err := h.personnelService.GetPersonnelByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "personnel not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, personnel)
}

func (h *PersonnelHandler) CreatePersonnel(c *gin.Context) {
	departmentPositionID, err := strconv.Atoi(c.PostForm("department_position_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid department_position_id"})
		return
	}

	req := models.PersonnelRequest{
		TypePersonnel:        c.PostForm("type_personnel"),
		DepartmentPositionID: departmentPositionID,
		ThaiAcademicPosition: strPtr(c.PostForm("thai_academic_position")),
		EngAcademicPosition:  strPtr(c.PostForm("eng_academic_position")),
		ThaiName:             c.PostForm("thai_name"),
		EngName:              c.PostForm("eng_name"),
		Education:            strPtr(c.PostForm("education")),
		RelatedFields:        strPtr(c.PostForm("related_fields")),
		Email:                strPtr(c.PostForm("email")),
		Website:              strPtr(c.PostForm("website")),
		AcademicPositionID:   intPtr(c.PostForm("academic_position_id")),
	}

	fileImage, _ := c.FormFile("file_image")

	createdPersonnel, err := h.personnelService.CreatePersonnel(req, fileImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPersonnel)
}

func (h *PersonnelHandler) UpdatePersonnel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid personnel ID"})
		return
	}

	departmentPositionID, err := strconv.Atoi(c.PostForm("department_position_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid department_position_id"})
		return
	}

	req := models.PersonnelRequest{
		TypePersonnel:        c.PostForm("type_personnel"),
		DepartmentPositionID: departmentPositionID,
		ThaiAcademicPosition: strPtr(c.PostForm("thai_academic_position")),
		EngAcademicPosition:  strPtr(c.PostForm("eng_academic_position")),
		ThaiName:             c.PostForm("thai_name"),
		EngName:              c.PostForm("eng_name"),
		Education:            strPtr(c.PostForm("education")),
		RelatedFields:        strPtr(c.PostForm("related_fields")),
		Email:                strPtr(c.PostForm("email")),
		Website:              strPtr(c.PostForm("website")),
		AcademicPositionID:   intPtr(c.PostForm("academic_position_id")),
	}

	fileImage, _ := c.FormFile("file_image")

	updatedPersonnel, err := h.personnelService.UpdatePersonnel(id, req, fileImage)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "personnel ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, updatedPersonnel)
}

func (h *PersonnelHandler) DeletePersonnel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "personnel ID required"})
		return
	}

	err = h.personnelService.DeletePersonnel(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "personnel not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "personnel deleted successfully"})
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPtr(s string) *int {
	if s == "" {
		return nil
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &val
}
