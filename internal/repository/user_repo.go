package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
)

type UserRepo interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(user *entity.User) error
	UpdateUserRole(userID string, userRoleID uint) error
	UpdateUserStatus(userID string, userStatusID uint) error
	DeleteUser(id string) error
}
