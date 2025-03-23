package entities

type UserStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}
