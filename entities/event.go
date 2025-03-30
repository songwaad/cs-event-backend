package entities

import "time"

type Event struct {
	EventID   uint       `json:"event_id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`

	Name              string    `json:"name" gorm:"unique"`
	Year              int       `json:"year"`
	Rationale         string    `json:"rationale"`
	Objective         string    `json:"objective"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	Location          string    `json:"location"`
	Methodology       string    `json:"methodology"`
	HasBudget         bool      `json:"has_budget"`
	Monitoring        string    `json:"monitoring"`
	Product           string
	ProductIndicators string `json:"product_indicators"`
	ProductTarget     string `json:"product_target"`
	Result            string
	ResultIndicators  string `json:"result_indicators"`
	ResultTarget      string `json:"result_target"`

	// Foreign key
	EventTypeStatusID uint `json:"event_type_status_id"`
	EventTypeStatus   EventTypeStatus
	EventPlanID       uint `json:"event_plan_id"`
	EventPlan         EventPlan
	EventTypeID       uint `json:"event_type_id"`
	EventType         EventType
	EventStrategyID   uint `json:"event_strategy_id"`
	EventStrategy     EventStrategy
	CreatedByUserID   string `json:"created_by_user_id"`
	User              User   `gorm:"foreignKey:CreatedByUserID;references:UserID"`
	EventStatusID     uint   `json:"event_status_id"`
	EventStatus       EventStatus

	// Many-to-many relationships
	Speakers         []Speaker `gorm:"many2many:event_speaker;"`
	ResponsibleUsers []User    `gorm:"many2many:event_responsible;"`
}
