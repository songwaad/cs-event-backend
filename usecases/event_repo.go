package usecases

import (
	"github.com/songwaad/cs-event-backend/entities"
)

// import (
// 	"github.com/songwaad/cs-event-backend/entities"
// )

type EventRepo interface {
	GetAllEventTypeStatus() ([]entities.EventTypeStatus, error)
	GetAllEventType() ([]entities.EventType, error)
	GetAllEventStatus() ([]entities.EventStatus, error)
	GetAllEventPlan() ([]entities.EventPlan, error)

	CreateEvent(event *entities.Event) error
	GetEventByID(eventID uint) (*entities.Event, error)
	GetAllEvents() ([]entities.Event, error)
	UpdateEvent(event *entities.Event) error
	DeleteEvent(eventID uint) error

	AddSpeakersToEvent(event *entities.Event, speakerIDs []uint) error
	AddResponsibleUsersToEvent(event *entities.Event, userIDs []string) error
}
