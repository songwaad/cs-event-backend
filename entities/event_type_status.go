package entities

type EventTypeStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}
