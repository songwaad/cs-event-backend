package entities

import "time"

type Event struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index"`

	Year           int `json:"year"`
	EventDetailsID uint
	EventDetails   EventDetails
	CreatedBy      string `json:"created_by"`
	User           User   `gorm:"foreignKey:CreatedBy"`
}
