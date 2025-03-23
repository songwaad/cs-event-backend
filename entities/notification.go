package entities

type Notification struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	Active         bool `json:"active"`
	UserId         User
	User           User
	EventDetailsID EventDetails
	EventDetails   EventDetails
}
