package entities

import (
	"time"
)

type User struct {
	ID           string `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     *time.Time `gorm:"index"`
	Email        string     `gorm:"unique"`
	Password     string
	FirstName    string `json:"firstname"`
	Lastname     string
	UserRoleID   uint
	UserRole     UserRole
	UserStatusID uint
	UserStatus   UserStatus
	Status       string
	EventDetails []EventDetails `gorm:"many2many:event_responsible;"`
}
