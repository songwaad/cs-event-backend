package dto

import (
	"time"

	"github.com/songwaad/cs-event-backend/entities"
)

type EventCreateDTO struct {
	Name        string    `json:"name" gorm:"unique"`
	Year        int       `json:"year"`
	Rationale   string    `json:"rationale"`
	Objective   string    `json:"objective"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Location    string    `json:"location"`
	Methodology string    `json:"methodology"`
	HasBudget   bool      `json:"has_budget"`
	Monitoring  string    `json:"monitoring"`

	// Foreign key
	EventTypeStatusID uint   `json:"event_type_status_id"`
	EventPlanID       uint   `json:"event_plan_id"`
	EventTypeID       uint   `json:"event_type_id"`
	EventStrategyID   uint   `json:"event_strategy_id"`
	CreatedByUserID   string `json:"created_by_user_id"`
	EventStatusID     uint   `json:"event_status_id"`

	// Many-to-many relationships
	ResponsibleUserIDs []string `json:"responsible_user_ids"`
	SpeakerIDs         []uint   `json:"speaker_ids"`
}

type EventResponseDTO struct {
	EventID     uint      `json:"event_id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique"`
	Year        int       `json:"year"`
	Rationale   string    `json:"rationale"`
	Objective   string    `json:"objective"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Location    string    `json:"location"`
	Methodology string    `json:"methodology"`
	HasBudget   bool      `json:"has_budget"`
	Monitoring  string    `json:"monitoring"`

	// Foreign key
	EventTypeStatusID uint `json:"event_type_status_id"`
	TypeStatus        string
	EventPlanID       uint `json:"event_plan_id"`
	Plan              string
	WorkPlan          string
	EventTypeID       uint `json:"event_type_id"`
	Type              string
	EventStrategyID   uint `json:"event_strategy_id"`
	Strategy          string
	Goal              string
	Tactic            string
	CreatedByUserID   string `json:"created_by_user_id"`
	FirstName         string
	Lastname          string
	EventStatusID     uint `json:"event_status_id"`
	Status            string

	// Many-to-many relationships
	ResponsibleUsers []UserResponseDTO `json:"responsible_users"`
	Speakers         []SpeakerDTO      `json:"speakers"`
}

func ToEventResponseDTO(event *entities.Event) *EventResponseDTO {
	var responsibleUsers []UserResponseDTO
	for _, user := range event.ResponsibleUsers {
		responsibleUsers = append(responsibleUsers, UserResponseDTO{
			UserID:    user.UserID,
			FirstName: user.FirstName,
			Lastname:  user.Lastname,
		})
	}

	var speakers []SpeakerDTO
	for _, speaker := range event.Speakers {
		speakers = append(speakers, SpeakerDTO{
			SpeakerID:   speaker.SpeakerID,
			FirstName:   speaker.FirstName,
			Lastname:    speaker.Lastname,
			Description: speaker.Description,
		})
	}

	return &EventResponseDTO{
		EventID:           event.EventID,
		Name:              event.Name,
		Year:              event.Year,
		Rationale:         event.Rationale,
		Objective:         event.Objective,
		StartDate:         event.StartDate,
		EndDate:           event.EndDate,
		Location:          event.Location,
		Methodology:       event.Methodology,
		HasBudget:         event.HasBudget,
		Monitoring:        event.Monitoring,
		EventTypeStatusID: event.EventTypeStatusID,
		TypeStatus:        event.EventTypeStatus.Status,
		EventPlanID:       event.EventPlanID,
		Plan:              event.EventPlan.WorkPlan,
		WorkPlan:          event.EventPlan.Work,
		EventTypeID:       event.EventTypeID,
		Type:              event.EventType.Type,
		EventStrategyID:   event.EventStrategyID,
		Strategy:          event.EventStrategy.Strategy.Strategy,
		Goal:              event.EventStrategy.Goal,
		Tactic:            event.EventStrategy.Tactic,
		CreatedByUserID:   event.CreatedByUserID,
		FirstName:         event.User.FirstName,
		Status:            event.EventStatus.Status,
		ResponsibleUsers:  responsibleUsers,
		Speakers:          speakers,
		Lastname:          event.User.Lastname,
	}
}
