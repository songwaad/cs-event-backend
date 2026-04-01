package database

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entity.UserRole{},
		&entity.UserStatus{},
		&entity.User{},
		&entity.Speaker{},
		&entity.EventPlan{},
		&entity.EventType{},
		&entity.EventStatus{},
		&entity.EventTypeStatus{},
		&entity.Strategy{},
		&entity.EventStrategy{},
		&entity.Event{},
		&entity.Budget{},
		&entity.Notification{},
	)
}
