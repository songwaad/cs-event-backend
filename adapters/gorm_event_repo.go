package adapters

import (
	"errors"
	"fmt"
	"time"

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

func (r *GormEventRepo) Create(event *entities.Event) error {
	return r.DB.Debug().Transaction(func(tx *gorm.DB) error {
		// บันทึก Result ก่อน
		if err := tx.Create(&event.EventDetails.EventResult).Error; err != nil {
			return err
		}

		// กำหนด EventResultID ให้ EventDetails
		event.EventDetails.EventResultID = event.EventDetails.EventResult.ID

		// บันทึก EventDetails ก่อน
		if err := tx.Create(&event.EventDetails).Error; err != nil {
			return err
		}

		// กำหนด EventDetailsID ให้ Event
		event.EventDetailsID = event.EventDetails.ID

		// ดึง ResponsibleUsers จาก DB
		if len(event.EventDetails.ResponsibleUsers) > 0 {
			var users []entities.User
			userIDs := make([]string, len(event.EventDetails.ResponsibleUsers))
			for i, user := range event.EventDetails.ResponsibleUsers {
				userIDs[i] = user.ID
			}
			if err := tx.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
				return err
			}
			event.EventDetails.ResponsibleUsers = users
		}

		// ดึง Instructor จาก DB
		if len(event.EventDetails.Instructor) > 0 {
			var instructors []entities.Instructor
			instructorIDs := make([]uint, len(event.EventDetails.Instructor))
			for i, instr := range event.EventDetails.Instructor {
				instructorIDs[i] = instr.ID
			}
			if err := tx.Where("id IN ?", instructorIDs).Find(&instructors).Error; err != nil {
				return err
			}
			event.EventDetails.Instructor = instructors
		}

		// บันทึก Event
		if err := tx.Create(event).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *GormEventRepo) GetByID(id int) (*entities.Event, error) {
	var event entities.Event
	result := r.DB.
		Preload("EventDetails.EventTypeStatus").
		Preload("EventDetails.EventPlane").
		Preload("EventDetails.EventType").
		Preload("EventDetails.EventStrategy").
		Preload("EventDetails.EventStrategy.Strategy").
		Preload("EventDetails.Instructor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "lastname", "description")
		}).
		Preload("EventDetails.ResponsibleUsers", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "lastname")
		}).
		Preload("EventDetails.EventResult").
		Preload("EventStatus").
		Preload("EventDetails").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "lastname")
		}).
		First(&event, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &event, nil
}

func (r *GormEventRepo) GetAll() ([]entities.Event, error) {
	var events []entities.Event
	result := r.DB.
		Preload("EventDetails.EventTypeStatus").
		Preload("EventDetails.EventPlane").
		Preload("EventDetails.EventType").
		Preload("EventDetails.EventStrategy").
		Preload("EventDetails.EventStrategy.Strategy").
		Preload("EventDetails.Instructor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "lastname", "description")
		}).
		Preload("EventDetails.ResponsibleUsers", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "lastname")
		}).
		Preload("EventDetails.EventResult").
		Preload("EventStatus").
		Preload("EventDetails").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "lastname")
		}).
		Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}

func (r *GormEventRepo) Update(event *entities.Event) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// ตรวจสอบว่า Event มีอยู่ใน DB หรือไม่
		var existingEvent entities.Event
		if err := tx.First(&existingEvent, event.ID).Error; err != nil {
			return fmt.Errorf("event with ID %d not found", event.ID)
		}

		// ตรวจสอบ Unique Constraint ของ EventDetails.Name
		if event.EventDetails.Name != "" {
			var conflictingEventDetails entities.EventDetails
			if err := tx.Where("name = ? AND id != ?", event.EventDetails.Name, event.EventDetails.ID).
				First(&conflictingEventDetails).Error; err == nil {
				return fmt.Errorf("event details with name '%s' already exists", event.EventDetails.Name)
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		}

		// ดึง ResponsibleUsers จาก DB
		if len(event.EventDetails.ResponsibleUsers) > 0 {
			var users []entities.User
			userIDs := make([]string, len(event.EventDetails.ResponsibleUsers))
			for i, user := range event.EventDetails.ResponsibleUsers {
				userIDs[i] = user.ID
			}
			if err := tx.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
				return err
			}
			event.EventDetails.ResponsibleUsers = users
		}

		// ดึง Instructor จาก DB
		if len(event.EventDetails.Instructor) > 0 {
			var instructors []entities.Instructor
			instructorIDs := make([]uint, len(event.EventDetails.Instructor))
			for i, instr := range event.EventDetails.Instructor {
				instructorIDs[i] = instr.ID
			}
			if err := tx.Where("id IN ?", instructorIDs).Find(&instructors).Error; err != nil {
				return err
			}
			event.EventDetails.Instructor = instructors
		}

		// อัปเดต Event และความสัมพันธ์ทั้งหมด
		if err := tx.Save(event).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *GormEventRepo) Delete(id int) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// หา Event ที่จะลบ
		var event entities.Event
		if err := tx.First(&event, id).Error; err != nil {
			return fmt.Errorf("event with ID %d not found", id)
		}

		// Soft Delete EventResult
		if err := tx.Model(&event.EventDetails.EventResult).Update("DeletedAt", time.Now()).Error; err != nil {
			return err
		}

		// Soft Delete EventDetails
		if err := tx.Model(&event.EventDetails).Update("DeletedAt", time.Now()).Error; err != nil {
			return err
		}

		// Soft Delete Event
		if err := tx.Delete(&event).Error; err != nil {
			return err
		}

		return nil
	})
}
