package usecases

import "github.com/songwaad/cs-event-backend/entities"

// import "github.com/songwaad/cs-event-backend/entities"

type EventUseCase interface {
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

type EventService struct {
	repo EventRepo
}

func NewEventService(repo EventRepo) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) GetAllEventTypeStatus() ([]entities.EventTypeStatus, error) {
	return s.repo.GetAllEventTypeStatus()
}
func (s *EventService) GetAllEventType() ([]entities.EventType, error) {
	return s.repo.GetAllEventType()
}

func (s *EventService) GetAllEventStatus() ([]entities.EventStatus, error) {
	return s.repo.GetAllEventStatus()
}

func (s *EventService) GetAllEventPlan() ([]entities.EventPlan, error) {
	return s.repo.GetAllEventPlan()
}

func (s *EventService) CreateEvent(event *entities.Event) error {
	return s.repo.CreateEvent(event)
}

func (s *EventService) GetEventByID(eventID uint) (*entities.Event, error) {
	return s.repo.GetEventByID(eventID)
}

func (s *EventService) GetAllEvents() ([]entities.Event, error) {
	return s.repo.GetAllEvents()
}

func (s *EventService) UpdateEvent(event *entities.Event) error {
	return s.repo.UpdateEvent(event)
}

func (s *EventService) DeleteEvent(eventID uint) error {
	return s.repo.DeleteEvent(eventID)
}

func (s *EventService) AddResponsibleUsersToEvent(event *entities.Event, userIDs []string) error {
	return s.repo.AddResponsibleUsersToEvent(event, userIDs)
}

func (s *EventService) AddSpeakersToEvent(event *entities.Event, speakerIDs []uint) error {
	return s.repo.AddSpeakersToEvent(event, speakerIDs)
}
