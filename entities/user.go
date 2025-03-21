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
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
	Email     string         `gorm:"unique"`
	Password  string
	FirstName string
	Lastname  string
	Role      string
	Status    string
}
