package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
)

type EventRepo interface {
	GetAllEventTypeStatus() ([]entity.EventTypeStatus, error)
	GetAllEventType() ([]entity.EventType, error)
	GetAllEventStatus() ([]entity.EventStatus, error)
	GetAllEventPlan() ([]entity.EventPlan, error)

	CreateEvent(event *entity.Event) error
	GetEventByID(eventID uint) (*entity.Event, error)
	GetAllEvents() ([]entity.Event, error)
	UpdateEvent(event *entity.Event) error
	DeleteEvent(eventID uint) error

	AddSpeakersToEvent(event *entity.Event, speakerIDs []uint) error
	AddResponsibleUsersToEvent(event *entity.Event, userIDs []string) error
}
