package entities

type Budget struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	EventDetailsID uint
	EventDetails   EventDetails
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
}
