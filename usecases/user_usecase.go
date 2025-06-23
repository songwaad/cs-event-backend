package usecases

import (
	"errors"

	"github.com/songwaad/cs-event-backend/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(user *entities.User) error
	Login(email, password string) (*entities.User, error)
	GetUserByID(userID string) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	DeleteUser(userID string) error
	ChangePassword(userEmail string, oldPassword string, newPassword string) error
	UpdateUserRole(userID string, userRoleID uint) error
	UpdateUserStatus(userID string, userStatusID uint) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *entities.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)
}

func (s *UserService) Login(email, password string) (*entities.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(userID string) (*entities.User, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	return s.repo.GetUserByID(userID)
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("Old password is incorrect")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repo.UpdateUser(user)
}

func (s *UserService) UpdateUserRole(userID string, userRoleID uint) error {
	return s.repo.UpdateUserRole(userID, userRoleID)
}

func (s *UserService) UpdateUserStatus(userID string, userStatusID uint) error {
	return s.repo.UpdateUserStatus(userID, userStatusID)
}
