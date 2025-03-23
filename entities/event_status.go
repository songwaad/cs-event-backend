package entities

type EventStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}
