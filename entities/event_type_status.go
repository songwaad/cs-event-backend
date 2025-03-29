package entities

type EventTypeStatus struct {
	EventTypeStatusID uint   `json:"event_type_status_id" gorm:"primaryKey"`
	Status            string `json:"status"`
}
