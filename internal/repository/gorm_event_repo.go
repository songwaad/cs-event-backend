package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"gorm.io/gorm"
)

type GormEventRepo struct {
	DB *gorm.DB
}

func NewGormEventRepo(DB *gorm.DB) EventRepo {
	return &GormEventRepo{DB: DB}
}

func (r *GormEventRepo) GetAllEventTypeStatus() ([]entity.EventTypeStatus, error) {
	var eventTypeStatus []entity.EventTypeStatus
	result := r.DB.Find(&eventTypeStatus)
	if result.Error != nil {
		return nil, result.Error
	}
	return eventTypeStatus, nil
}

func (r *GormEventRepo) GetAllEventType() ([]entity.EventType, error) {
	var eventTypes []entity.EventType
	result := r.DB.Find(&eventTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return eventTypes, nil
}

func (r *GormEventRepo) GetAllEventStatus() ([]entity.EventStatus, error) {
	var eventStatus []entity.EventStatus
	result := r.DB.Find(&eventStatus)
	if result.Error != nil {
		return nil, result.Error
	}
	return eventStatus, nil
}

func (r *GormEventRepo) GetAllEventPlan() ([]entity.EventPlan, error) {
	var EventPlans []entity.EventPlan
	result := r.DB.Find(&EventPlans)
	if result.Error != nil {
		return nil, result.Error
	}
	return EventPlans, nil
}

func (r *GormEventRepo) CreateEvent(event *entity.Event) error {
	return r.DB.Create(event).Error
}

func (r *GormEventRepo) GetEventByID(eventID uint) (*entity.Event, error) {
	var event entity.Event
	err := r.DB.
		Preload("EventTypeStatus").
		Preload("EventPlan").
		Preload("EventType").
		Preload("EventStrategy").
		Preload("EventStrategy.Strategy").
		Preload("EventStatus").
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

func (r *GormEventRepo) GetAllEvents() ([]entity.Event, error) {
	var events []entity.Event
	err := r.DB.
		Preload("EventTypeStatus").
		Preload("EventPlan").
		Preload("EventType").
		Preload("EventStrategy").
		Preload("EventStrategy.Strategy").
		Preload("EventStatus").
		Preload("ResponsibleUsers").
		Preload("Speakers").
		Preload("User").
		Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *GormEventRepo) UpdateEvent(event *entity.Event) error {
	return r.DB.Model(&entity.Event{}).Where("event_id = ?", event.EventID).Save(*event).Error
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

	return r.DB.Delete(&entity.Event{}, eventID).Error
}

func (r *GormEventRepo) AddResponsibleUsersToEvent(event *entity.Event, userIDs []string) error {
	for _, userID := range userIDs {
		var user entity.User
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

func (r *GormEventRepo) AddSpeakersToEvent(event *entity.Event, speakerIDs []uint) error {
	for _, speakerID := range speakerIDs {
		var speaker entity.Speaker
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
