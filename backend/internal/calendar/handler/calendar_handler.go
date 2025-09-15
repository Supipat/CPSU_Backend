package handler

import (
	"net/http"
	"time"

	"cpsu/internal/calendar/service"

	"github.com/gin-gonic/gin"
)

type CalendarHandler struct {
	calendarService *service.CalendarService
}

func NewCalendarHandler(calendarService *service.CalendarService) *CalendarHandler {
	return &CalendarHandler{calendarService: calendarService}
}

func (h *CalendarHandler) GetAllCalendar(c *gin.Context) {
	calendarID := c.Query("calendarId")

	now := time.Now().UTC()
	timeMin := now.Add(-7 * 24 * time.Hour)
	timeMax := now.Add(30 * 24 * time.Hour)

	events, err := h.calendarService.GetAllCalendar(c.Request.Context(), calendarID, timeMin, timeMax, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(events),
		"items": events,
	})
}
