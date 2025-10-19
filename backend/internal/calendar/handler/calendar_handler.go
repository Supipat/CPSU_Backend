package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"cpsu/internal/calendar/models"
	"cpsu/internal/calendar/service"

	"github.com/gin-gonic/gin"
)

type CalendarHandler struct {
	calendarService service.CalendarService
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

	c.JSON(http.StatusOK, gin.H{"message": "calendar deleted successfully"})
}
