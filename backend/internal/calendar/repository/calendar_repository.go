package repository

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"cpsu/internal/calendar/models"

	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type CalendarRepository interface {
	GetAllCalendar(ctx context.Context, calendarID string, timeMin, timeMax time.Time, maxResults int) ([]models.Calendar, error)
}

type calendarRepo struct {
	svc *calendar.Service
}

func NewCalendarRepository(ctx context.Context) (CalendarRepository, error) {
	b64 := os.Getenv("GOOGLE_CRED_JSON_BASE64")
	if b64 == "" {
		return nil, fmt.Errorf("GOOGLE_CRED_JSON_BASE64 is not set")
	}

	bytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("decode credentials: %w", err)
	}

	cfg, err := google.JWTConfigFromJSON(bytes, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("parse jwt config: %w", err)
	}

	httpClient := cfg.Client(ctx)
	svc, err := calendar.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("new calendar service: %w", err)
	}

	return &calendarRepo{svc: svc}, nil
}

func (r *calendarRepo) GetAllCalendar(ctx context.Context, calendarID string, timeMin, timeMax time.Time, maxResults int) ([]models.Calendar, error) {
	if calendarID == "" {
		calendarID = "primary"
	}

	call := r.svc.Events.List(calendarID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(timeMin.Format(time.RFC3339)).
		TimeMax(timeMax.Format(time.RFC3339)).
		OrderBy("startTime")

	if maxResults > 0 {
		call = call.MaxResults(int64(maxResults))
	}

	res, err := call.Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	events := make([]models.Calendar, 0, len(res.Items))
	for _, e := range res.Items {
		start := e.Start.DateTime
		if start == "" {
			start = e.Start.Date
		}
		end := e.End.DateTime
		if end == "" {
			end = e.End.Date
		}

		events = append(events, models.Calendar{
			ID:          e.Id,
			Title:       e.Summary,
			Description: e.Description,
			Start:       start,
			End:         end,
			Location:    e.Location,
			HtmlLink:    e.HtmlLink,
		})
	}
	return events, nil
}
