package entities

import (
	"time"
)

type EventDetails struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	Name              string     `json:"name" gorm:"unique"`
	DeleteAt          *time.Time `gorm:"index"`
	EventTypeStatusID uint
	EventTypeStatus   EventTypeStatus
	EventPlaneID      uint
	EventPlane        EventPlane
	EventTypeID       uint
	EventType         EventType
	Rationale         string `json:"rationale"`
	EventStrategyID   uint
	EventStrategy     EventStrategy
	Objective         string    `json:"objective"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	Location          string    `json:"location"`
	Methodology       string    `json:"methodology"`
	HasBudget         bool      `json:"has_budget"`
	Monitoring        string    `json:"monitoring"`
	EventResultID     uint
	EventResult       EventResult
	Instructor        []Instructor `gorm:"many2many:event_instructor;"`
	ResponsibleUsers  []User       `gorm:"many2many:event_responsible;"`
}
