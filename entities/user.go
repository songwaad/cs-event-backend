package entities

import (
	"time"

	"gorm.io/gorm"
)

type Role int

const (
	Admin Role = iota
	Staff
	Professor
	Student
)

type User struct {
	ID           string `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     gorm.DeletedAt `gorm:"index"`
	Email        string         `gorm:"unique"`
	Password     string
	FirstName    string
	Lastname     string
	UserRoleID   uint
	UserRole     UserRole
	UserStatusID uint
	UserStatus   UserStatus
	Status       string
	// EventDetails []EventDetails `gorm:"many2many:event_reponsible;"`
}
