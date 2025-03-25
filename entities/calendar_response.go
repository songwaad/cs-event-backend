package entities

import "time"

type CalendarResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Location  string    `json:"location"`
	EventType string    `json:"eventType"`
}
