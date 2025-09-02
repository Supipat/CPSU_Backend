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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personnels, err := h.personnelService.GetAllPersonnels(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, personnels)
}

func (h *PersonnelHandler) GetPersonnelByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID required"})
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
	Ptr := func(s string) *string {
		if s == "" {
			return nil
		}
		return &s
	}

	IntPtr := func(s string) *int {
		if s == "" {
			return nil
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &val
	}

	typeID, err := strconv.Atoi(c.PostForm("type_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type_id"})
		return
	}

	departmentPositionID, err := strconv.Atoi(c.PostForm("department_position_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid department_position_id"})
		return
	}

	req := models.PersonnelRequest{
		TypeID:               typeID,
		DepartmentPositionID: departmentPositionID,
		AcademicPositionID:   IntPtr(c.PostForm("academic_position_id")),
		ThaiName:             c.PostForm("thai_name"),
		EngName:              c.PostForm("eng_name"),
		Education:            Ptr(c.PostForm("education")),
		RelatedFields:        Ptr(c.PostForm("related_fields")),
		Email:                Ptr(c.PostForm("email")),
		Website:              Ptr(c.PostForm("website")),
	}

	createdPersonnel, err := h.personnelService.CreatePersonnel(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPersonnel)
}

func (h *PersonnelHandler) UpdatePersonnel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid personnel ID"})
		return
	}

	Ptr := func(s string) *string {
		if s == "" {
			return nil
		}
		return &s
	}

	IntPtr := func(s string) *int {
		if s == "" {
			return nil
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &val
	}

	typeID, err := strconv.Atoi(c.PostForm("type_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type_id"})
		return
	}

	departmentPositionID, err := strconv.Atoi(c.PostForm("department_position_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid department_position_id"})
		return
	}

	req := models.PersonnelRequest{
		TypeID:               typeID,
		DepartmentPositionID: departmentPositionID,
		AcademicPositionID:   IntPtr(c.PostForm("academic_position_id")),
		ThaiName:             c.PostForm("thai_name"),
		EngName:              c.PostForm("eng_name"),
		Education:            Ptr(c.PostForm("education")),
		RelatedFields:        Ptr(c.PostForm("related_fields")),
		Email:                Ptr(c.PostForm("email")),
		Website:              Ptr(c.PostForm("website")),
	}

	updatedPersonnel, err := h.personnelService.UpdatePersonnel(id, req)
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
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
