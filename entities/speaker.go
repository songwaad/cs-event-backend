package entities

import (
	"time"
)

type Speaker struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeleteAt     *time.Time     `gorm:"index"`
	FirstName    string         `json:"firstname"`
	Lastname     string         `json:"lastname"`
	Description  string         `json:"description"`
	EventDetails []EventDetails `gorm:"many2many:event_instructor;"`
}
