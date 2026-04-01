package usecase

import (
	"errors"

	"github.com/songwaad/cs-event-backend/internal/entity"
	"github.com/songwaad/cs-event-backend/internal/repository"
	"github.com/songwaad/cs-event-backend/pkg/hash"
)

type UserUseCase interface {
	Register(user *entity.User) error
	Login(email, password string) (*entity.User, error)
	GetUserByID(userID string) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
	DeleteUser(userID string) error
	ChangePassword(userEmail string, oldPassword string, newPassword string) error
	UpdateUserRole(userID string, userRoleID uint) error
	UpdateUserStatus(userID string, userStatusID uint) error
}

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *entity.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already in use")
	}

	hashedPassword, err := hash.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s *UserService) Login(email, password string) (*entity.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := hash.Compare(user.Password, password); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(userID string) (*entity.User, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	return s.repo.GetUserByID(userID)
}

func (s *UserService) GetAllUsers() ([]entity.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) DeleteUser(userID string) error {
	if userID == "" {
		return errors.New("user ID is required")
	}

	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(userID)
}

func (s *UserService) ChangePassword(userEmail string, oldPassword string, newPassword string) error {
	user, err := s.repo.GetUserByEmail(userEmail)
	if err != nil {
		return errors.New("user not found")
	}

	if err := hash.Compare(user.Password, oldPassword); err != nil {
		return errors.New("Old password is incorrect")
	}

	hashedPassword, err := hash.Hash(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.repo.UpdateUser(user)
}

func (s *UserService) UpdateUserRole(userID string, userRoleID uint) error {
	return s.repo.UpdateUserRole(userID, userRoleID)
}

func (s *UserService) UpdateUserStatus(userID string, userStatusID uint) error {
	return s.repo.UpdateUserStatus(userID, userStatusID)
}
