package entities

import (
	"time"
)

type Instructor struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	FirstName   string     `json:"firstname"`
	Lastname    string     `json:"lastname"`
	Description string     `json:"description"`
}
