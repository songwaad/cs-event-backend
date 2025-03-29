package usecases

import (
	"github.com/songwaad/cs-event-backend/entities"
)

type UserRepo interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByID(id string) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(user *entities.User) error
	UpdateUserRole(userID string, userRoleID uint) error
	UpdateUserStatus(userID string, userStatusID uint) error
	DeleteUser(id string) error
}
