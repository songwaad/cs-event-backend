package usecases

import (
	"errors"

	"github.com/songwaad/cs-event-backend/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(user *entities.User) error
	Login(email, password string) (*entities.User, error)
	GetUserByID(id string) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id string) error
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

func (s *UserService) GetUserByID(id string) (*entities.User, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}

	return s.repo.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUser(user *entities.User) error {
	if user.ID == "" {
		return errors.New("user ID is required")
	}

	existingUser, err := s.repo.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	if user.Password != "" && user.Password != existingUser.Password {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	} else {
		user.Password = existingUser.Password
	}

	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("user ID is required")
	}

	_, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(id)
}
