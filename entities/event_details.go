package entities

import (
	"time"
)

type EventDetails struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Name              string `json:"name" gorm:"unique"`
	Year              int    `json:"year"`
	EventTypeStatusID EventTypeStatus
	EventTypeStatus   EventTypeStatus
	EventPlaneID      EventPlane
	EventPlane        EventPlane
	EventTypeID       EventType
	EventType         EventType
	Rationale         string `json:"retionale"`
	EventStrategyID   EventStrategy
	EventStrategy     EventStrategy
	EventGoalID       EventGoal
	EventGoal         EventGoal
	EventTacticID     EventTactic
	EventTactic       EventTactic
	Objective         string    `json:"objective"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	Location          string    `json:"location"`
	Methodology       string    `json:"methodology"`
	HasBudget         bool      `json:"has_budget"`
	Monitoring        string    `json:"monitoring"`
	EventStatusID     EventStatus
	EventStatus       EventStatus
	Instructor        []Instructor `gorm:"many2many:event_instructor;"`
	ResponsibleUsers  []User       `gorm:"many2many:event_reponsible;"`
}
