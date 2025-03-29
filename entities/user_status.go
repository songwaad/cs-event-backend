package entities

type UserStatus struct {
	UserStatusID uint   `json:"user_status_id" gorm:"primaryKey"`
	Status       string `json:"status"`
}
