package usecases

import "github.com/songwaad/cs-event-backend/entities"

type EventUseCase interface {
	CreateEvent(event *entities.Event) error
	GetEventByID(id int) (*entities.Event, error)
	GetAllEvents() ([]entities.Event, error)
	UpdateEvent(event *entities.Event) error
	DeleteEvent(id int) error
}

type EventService struct {
	repo EventRepo
}

func NewEventService(repo EventRepo) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event *entities.Event) error {
	return s.repo.Create(event)
}

func (s *EventService) GetEventByID(id int) (*entities.Event, error) {
	return s.repo.GetByID(id)
}

func (s *EventService) GetAllEvents() ([]entities.Event, error) {
	return s.repo.GetAll()
}

func (s *EventService) UpdateEvent(event *entities.Event) error {
	return s.repo.Update(event)
}

func (s *EventService) DeleteEvent(id int) error {
	return s.repo.Delete(id)
}
