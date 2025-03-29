package entities

type UserRole struct {
	UserRoleID uint   `json:"user_role_id" gorm:"primaryKey"`
	Role       string `json:"role"`
}
