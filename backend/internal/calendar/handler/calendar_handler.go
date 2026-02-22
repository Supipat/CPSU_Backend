package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/calendar/models"
	"cpsu/internal/calendar/service"

	"cpsu/internal/auth/repository"

	"github.com/gin-gonic/gin"
)

type CalendarHandler struct {
	calendarService service.CalendarService
	auditRepo       *repository.AuditRepository
}

func NewCalendarHandler(calendarService service.CalendarService) *CalendarHandler {
	return &CalendarHandler{calendarService: calendarService}
}

func (h *CalendarHandler) GetAllCalendars(c *gin.Context) {
	var param models.CalendarQueryParam
	if err := c.BindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	calendars, err := h.calendarService.GetAllCalendars(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, calendars)
}

func (h *CalendarHandler) GetCalendarByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid calendar ID"})
		return
	}

	calendar, err := h.calendarService.GetCalendarByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "calendar not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, calendar)
}

func (h *CalendarHandler) CreateCalendar(c *gin.Context) {
	var req models.CalendarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCalendar, err := h.calendarService.CreateCalendar(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "create", "calendar",
		strconv.Itoa(createdCalendar.CalenderID),
		map[string]interface{}{
			"title": createdCalendar.Title,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

	c.JSON(http.StatusCreated, createdCalendar)
}

func (h *CalendarHandler) UpdateCalendar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid calendar ID"})
		return
	}

	var req models.CalendarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCalendar, err := h.calendarService.UpdateCalendar(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "calendar not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "update", "calendar",
		strconv.Itoa(updatedCalendar.CalenderID),
		map[string]interface{}{
			"title": updatedCalendar.Title,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

	c.JSON(http.StatusOK, updatedCalendar)
}

func (h *CalendarHandler) DeleteCalendar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid calendar ID"})
		return
	}

	err = h.calendarService.DeleteCalendar(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "calendar not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	userID := c.GetInt("user_id")

	_ = h.auditRepo.LogAudit(
		userID, "delete", "calendar", strconv.Itoa(id),
		map[string]interface{}{
			"calendar_id": id,
		},
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)

	c.JSON(http.StatusOK, gin.H{"message": "calendar deleted successfully"})
}
