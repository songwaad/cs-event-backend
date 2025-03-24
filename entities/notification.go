package entities

type Notification struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	Active         bool `json:"active"`
	UserId         string
	User           User
	EventDetailsID uint
	EventDetails   EventDetails
}
