package entities

type EventType struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Type string `json:"type"`
}
