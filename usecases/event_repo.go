package usecases

import "github.com/songwaad/cs-event-backend/entities"

// import (
// 	"github.com/songwaad/cs-event-backend/entities"
// )

type EventRepo interface {
	GetAllEventTypeStatus() ([]entities.EventTypeStatus, error)
	GetAllEventType() ([]entities.EventType, error)
	GetAllEventStatus() ([]entities.EventStatus, error)
	GetAllEventPlan() ([]entities.EventPlan, error)
	// Create(event *entities.Event) error
	// GetByID(id int) (*entities.Event, error)
	// GetAll() ([]entities.Event, error)
	// Update(event *entities.Event) error
	// Delete(id int) error
	// GetCalendarEvents() ([]entities.CalendarResponse, error)
}
