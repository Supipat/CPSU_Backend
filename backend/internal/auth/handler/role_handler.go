package handler

import (
	"net/http"
	"strconv"

	"cpsu/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	RoleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{RoleService: roleService}
}

func (h *RoleHandler) AssignRole(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user_id"})
		return
	}

	var req struct {
		RoleID int `json:"role_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	assignedBy := c.GetInt("user_id")

	if err := h.RoleService.AssignRole(userID, req.RoleID, assignedBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role assigned successfully"})
}
