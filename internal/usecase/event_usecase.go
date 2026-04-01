package usecase

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"github.com/songwaad/cs-event-backend/internal/repository"
)

type EventUseCase interface {
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

type EventService struct {
	repo repository.EventRepo
}

func NewEventService(repo repository.EventRepo) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) GetAllEventTypeStatus() ([]entity.EventTypeStatus, error) {
	return s.repo.GetAllEventTypeStatus()
}
func (s *EventService) GetAllEventType() ([]entity.EventType, error) {
	return s.repo.GetAllEventType()
}

func (s *EventService) GetAllEventStatus() ([]entity.EventStatus, error) {
	return s.repo.GetAllEventStatus()
}

func (s *EventService) GetAllEventPlan() ([]entity.EventPlan, error) {
	return s.repo.GetAllEventPlan()
}

func (s *EventService) CreateEvent(event *entity.Event) error {
	return s.repo.CreateEvent(event)
}

func (s *EventService) GetEventByID(eventID uint) (*entity.Event, error) {
	return s.repo.GetEventByID(eventID)
}

func (s *EventService) GetAllEvents() ([]entity.Event, error) {
	return s.repo.GetAllEvents()
}

func (s *EventService) UpdateEvent(event *entity.Event) error {
	return s.repo.UpdateEvent(event)
}

func (s *EventService) DeleteEvent(eventID uint) error {
	return s.repo.DeleteEvent(eventID)
}

func (s *EventService) AddResponsibleUsersToEvent(event *entity.Event, userIDs []string) error {
	return s.repo.AddResponsibleUsersToEvent(event, userIDs)
}

func (s *EventService) AddSpeakersToEvent(event *entity.Event, speakerIDs []uint) error {
	return s.repo.AddSpeakersToEvent(event, speakerIDs)
}
