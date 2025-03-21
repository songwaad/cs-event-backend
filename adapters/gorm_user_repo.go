package adapters

import (
	"github.com/songwaad/cs-event-backend/entities"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	DB *gorm.DB
}

func (r *GormUserRepo) CreateUser(user *entities.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepo) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepo) GetUserByID(userID string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepo) DeleteUser(userID string) error {
	return r.DB.Delete(&entities.User{}, userID).Error
}
