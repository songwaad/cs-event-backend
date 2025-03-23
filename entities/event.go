package entities

import "time"

type Event struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Year           int `json:"year"`
	EventDetailsID EventDetails
	EventDetails   EventDetails
	CreatedBy      string `json:"created_by"`
	User           User   `gorm:"foreignKey:CreatedBy"`
}
