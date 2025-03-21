package usecases

import (
	"github.com/songwaad/cs-event-backend/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepo UserRepo
}

func (u *UserUseCase) Register(user entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.UserRepo.CreateUser(&user)
}

func (u *UserUseCase) Login(email, password string) (*entities.User, error) {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) DeleteUser(id string) error {
	_, err := u.UserRepo.GetUserByID(id)
	if err != nil {
		return err
	}
	return u.UserRepo.DeleteUser(id)
}
