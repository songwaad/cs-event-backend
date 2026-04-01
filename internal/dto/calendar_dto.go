package dto

import (
	"time"

	"github.com/songwaad/cs-event-backend/internal/entity"
)

type CalendarResponseDTO struct {
	EventID   uint      `json:"event_id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Location  string    `json:"location"`
	EventType string    `json:"eventType"`
}

func ToCalendarResponseDTO(event *entity.Event) CalendarResponseDTO {
	return CalendarResponseDTO{
		EventID:   event.EventID,
		Name:      event.Name,
		StartDate: event.StartDate,
		EndDate:   event.EndDate,
		Location:  event.Location,
		EventType: event.EventTypeStatus.Status,
	}
}
