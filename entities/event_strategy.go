package entities

type EventStrategy struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Strategy string `json:"strategy"`
}
