package entities

import (
	"time"
)

type Speaker struct {
	SpeakerID uint `json:"speaker_id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	FirstName   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Description string `json:"description"`
	ImageUrl    string

	// Relationships
	Event []Event `gorm:"many2many:event_speaker;"`
}
