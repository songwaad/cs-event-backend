package entities

type EventStatus struct {
	EventStatusID uint   `json:"event_status_id" gorm:"primaryKey"`
	Status        string `json:"status"`
}
