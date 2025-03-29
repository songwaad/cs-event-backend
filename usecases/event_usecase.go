package usecases

import "github.com/songwaad/cs-event-backend/entities"

// import "github.com/songwaad/cs-event-backend/entities"

type EventUseCase interface {
	GetAllEventTypeStatus() ([]entities.EventTypeStatus, error)
	GetAllEventType() ([]entities.EventType, error)
	GetAllEventStatus() ([]entities.EventStatus, error)
	GetAllEventPlan() ([]entities.EventPlan, error)
	// 	CreateEvent(event *entities.Event) error
	// 	GetEventByID(id int) (*entities.Event, error)
	// 	GetAllEvents() ([]entities.Event, error)
	// 	UpdateEvent(event *entities.Event) error
	// 	DeleteEvent(id int) error
	// 	GetCalendarEvents() ([]entities.CalendarResponse, error)
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

// func (s *EventService) CreateEvent(event *entities.Event) error {
// 	return s.repo.Create(event)
// }

// func (s *EventService) GetEventByID(id int) (*entities.Event, error) {
// 	return s.repo.GetByID(id)
// }

// func (s *EventService) GetAllEvents() ([]entities.Event, error) {
// 	return s.repo.GetAll()
// }

// func (s *EventService) UpdateEvent(event *entities.Event) error {
// 	return s.repo.Update(event)
// }

// func (s *EventService) DeleteEvent(id int) error {
// 	return s.repo.Delete(id)
// }

// func (s *EventService) GetCalendarEvents() ([]entities.CalendarResponse, error) {
// 	return s.repo.GetCalendarEvents()
// }
