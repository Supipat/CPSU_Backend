package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/permission/models"
	"cpsu/internal/permission/service"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	userRoleService service.PermissionService
}

func NewPermissionHandler(userRoleService service.PermissionService) *PermissionHandler {
	return &PermissionHandler{userRoleService: userRoleService}
}

func (h *PermissionHandler) UpdateUserRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var req models.ChangeUserRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUserRole, err := h.userRoleService.UpdateUserRole(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, updatedUserRole)
}
