package entities

type Notification struct {
	NotifyID uint   `json:"id" gorm:"primaryKey"`
	Active   bool   `json:"active"`
	UserId   string `json:"user_id"`
	User     User
	EventID  uint `json:"event_id"`
	Event    Event
}
