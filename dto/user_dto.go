package dto

import (
	"time"

	"github.com/songwaad/cs-event-backend/entities"
)

type UserLoginDTO struct {
	Email    string
	Password string
}

type UserRegisterDTO struct {
	UserID    string     `json:"user_id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`

	Email     string `gorm:"unique"`
	Password  string
	FirstName string `json:"first_name"`
	Lastname  string

	// Foreign key
	UserRoleID   uint `json:"user_role_id"`
	UserStatusID uint `json:"user_status_id"`
}

type UserUpdateUserStatusDTO struct {
	UserStatusID uint `json:"user_status_id"`
}

type UserUpdateUserRoleDTO struct {
	UserRoleID uint `json:"user_role_id"`
}

type UserChangePasswordDTO struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserResponseDTO struct {
	UserID     string `json:"user_id" gorm:"primaryKey"`
	Email      string `gorm:"unique"`
	FirstName  string `json:"first_name"`
	Lastname   string `json:"last_name"`
	UserRole   UserRoleDTO
	UserStatus UserStatusDTO
}

type UserRoleDTO struct {
	UserRoleID uint   `json:"user_role_id"`
	Role       string `json:"role"`
}

type UserStatusDTO struct {
	UserStatusID uint   `json:"user_status_id"`
	Status       string `json:"status"`
}

func ToUserResponseDTO(entity entities.User) UserResponseDTO {
	return UserResponseDTO{
		UserID:     entity.UserID,
		Email:      entity.Email,
		FirstName:  entity.FirstName,
		Lastname:   entity.Lastname,
		UserRole:   ToUserRoleDTO(entity.UserRole),
		UserStatus: ToUserStatusDTO(entity.UserStatus),
	}
}

func ToUserRoleDTO(role entities.UserRole) UserRoleDTO {
	return UserRoleDTO{
		UserRoleID: role.UserRoleID,
		Role:       role.Role,
	}
}

func ToUserStatusDTO(role entities.UserStatus) UserStatusDTO {
	return UserStatusDTO{
		UserStatusID: role.UserStatusID,
		Status:       role.Status,
	}
}
