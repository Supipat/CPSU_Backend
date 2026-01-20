package handler

import (
	"net/http"

	"cpsu/internal/auth/models"
	"cpsu/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	var param models.UserQueryParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	users, err := h.UserService.GetAllUser(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
