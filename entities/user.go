package entities

import (
	"time"
)

type User struct {
	UserID    string `json:"user_id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	Email        string `gorm:"unique"`
	Password     string
	FirstName    string `json:"first_name"`
	Lastname     string
	UserRoleID   uint `json:"user_role_id"`
	UserRole     UserRole
	UserStatusID uint `json:"user_status_id"`
	UserStatus   UserStatus
	Status       string

	// Many-to-many relationships
	Event []Event `gorm:"many2many:event_responsible;"`
}
