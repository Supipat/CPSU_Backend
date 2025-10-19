package service

import (
	"cpsu/internal/calendar/models"
	"cpsu/internal/calendar/repository"
)

type CalendarService interface {
	GetAllCalendars(param models.CalendarQueryParam) ([]models.Calendar, error)
	GetCalendarByID(id int) (*models.Calendar, error)
	CreateCalendar(req models.CalendarRequest) (*models.Calendar, error)
	UpdateCalendar(id int, req models.CalendarRequest) (*models.Calendar, error)
	DeleteCalendar(id int) error
}

type calendarService struct {
	repo repository.CalendarRepository
}

func NewCalendarService(repo repository.CalendarRepository) CalendarService {
	return &calendarService{
		repo: repo,
	}
}

func (s *calendarService) GetAllCalendars(param models.CalendarQueryParam) ([]models.Calendar, error) {
	return s.repo.GetAllCalendars(param)
}

func (s *calendarService) GetCalendarByID(id int) (*models.Calendar, error) {
	return s.repo.GetCalendarByID(id)
}

func (s *calendarService) CreateCalendar(req models.CalendarRequest) (*models.Calendar, error) {
	return s.repo.CreateCalendar(&req)
}

func (s *calendarService) UpdateCalendar(id int, req models.CalendarRequest) (*models.Calendar, error) {
	req.CalenderID = id
	return s.repo.UpdateCalendar(&req)
}

func (s *calendarService) DeleteCalendar(id int) error {
	return s.repo.DeleteCalendar(id)
}
