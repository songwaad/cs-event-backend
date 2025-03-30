package adapters

import (
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
	"gorm.io/gorm"
)

type GormEventRepo struct {
	DB *gorm.DB
}

func NewGormEventRepo(DB *gorm.DB) usecases.EventRepo {
	return &GormEventRepo{DB: DB}
}

func (r *GormEventRepo) GetAllEventTypeStatus() ([]entities.EventTypeStatus, error) {
	var eventTypeStatus []entities.EventTypeStatus
	result := r.DB.Find(&eventTypeStatus)
	if result.Error != nil {
		return nil, result.Error
	}
	return eventTypeStatus, nil
}

func (r *GormEventRepo) GetAllEventType() ([]entities.EventType, error) {
	var eventTypes []entities.EventType
	result := r.DB.Find(&eventTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return eventTypes, nil
}

func (r *GormEventRepo) GetAllEventStatus() ([]entities.EventStatus, error) {
	var eventStatus []entities.EventStatus
	result := r.DB.Find(&eventStatus)
	if result.Error != nil {
		return nil, result.Error
	}
	return eventStatus, nil
}

func (r *GormEventRepo) GetAllEventPlan() ([]entities.EventPlan, error) {
	var EventPlans []entities.EventPlan
	result := r.DB.Find(&EventPlans)
	if result.Error != nil {
		return nil, result.Error
	}
	return EventPlans, nil
}

func (r *GormEventRepo) CreateEvent(event *entities.Event) error {
	return r.DB.Create(event).Error
}

func (r *GormEventRepo) GetEventByID(eventID uint) (*entities.Event, error) {
	var event entities.Event
	err := r.DB.
		Preload("EventTypeStatus").
		Preload("EventPlan").
		Preload("EventType").
		Preload("EventStrategy").
		Preload("EventStrategy.Strategy").
		Preload("ResponsibleUsers").
		Preload("Speakers").
		Preload("User").
		Where("event_id = ?", eventID).
		First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *GormEventRepo) GetAllEvents() ([]entities.Event, error) {
	var events []entities.Event
	err := r.DB.
		Preload("EventTypeStatus").
		Preload("EventPlan").
		Preload("EventType").
		Preload("EventStrategy").
		Preload("EventStrategy.Strategy").
		Preload("ResponsibleUsers").
		Preload("Speakers").
		Preload("User").
		Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *GormEventRepo) UpdateEvent(event *entities.Event) error {
	return r.DB.Model(&entities.Event{}).Where("event_id = ?", event.EventID).Save(*event).Error
}

func (r *GormEventRepo) DeleteEvent(eventID uint) error {
	// ลบความสัมพันธ์ในตารางกลาง event_responsible ก่อน
	if err := r.DB.Exec("DELETE FROM event_responsible WHERE event_event_id = ?", eventID).Error; err != nil {
		return err
	}

	// ลบความสัมพันธ์ในตารางกลาง event_speaker ก่อน
	if err := r.DB.Exec("DELETE FROM event_speaker WHERE event_event_id = ?", eventID).Error; err != nil {
		return err
	}

	return r.DB.Delete(&entities.Event{}, eventID).Error
}

func (r *GormEventRepo) AddResponsibleUsersToEvent(event *entities.Event, userIDs []string) error {
	for _, userID := range userIDs {
		var user entities.User
		if err := r.DB.First(&user, "user_id = ?", userID).Error; err != nil {
			return err
		}
		// เพิ่ม user ไปยัง event ผ่านความสัมพันธ์ many-to-many
		if err := r.DB.Model(event).Association("ResponsibleUsers").Append(&user); err != nil {
			return err
		}
	}
	return nil
}

func (r *GormEventRepo) AddSpeakersToEvent(event *entities.Event, speakerIDs []uint) error {
	for _, speakerID := range speakerIDs {
		var speaker entities.Speaker
		if err := r.DB.First(&speaker, speakerID).Error; err != nil {
			return err
		}
		// เพิ่ม speaker ไปยัง event
		if err := r.DB.Model(event).Association("Speakers").Append(&speaker); err != nil {
			return err
		}
	}
	return nil
}
