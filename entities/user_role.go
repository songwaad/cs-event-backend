package entities

type UserRole struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Role string `json:"role"`
}
