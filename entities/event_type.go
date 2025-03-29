package entities

type EventType struct {
	EventTypeID uint   `json:"event_type_id" gorm:"primaryKey"`
	Type        string `json:"type"`
}
