package service

import (
	"context"
	"time"

	"cpsu/internal/calendar/models"
	"cpsu/internal/calendar/repository"
)

type CalendarService struct {
	repo       repository.CalendarRepository
	defaultCal string
}

func NewCalendarService(repo repository.CalendarRepository, defaultCal string) *CalendarService {
	return &CalendarService{repo: repo, defaultCal: defaultCal}
}

func (s *CalendarService) GetAllCalendar(ctx context.Context, calendarID string, timeMin, timeMax time.Time, maxResults int) ([]models.Calendar, error) {
	if calendarID == "" {
		calendarID = s.defaultCal
	}
	return s.repo.GetAllCalendar(ctx, calendarID, timeMin, timeMax, maxResults)
}
