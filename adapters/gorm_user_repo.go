package adapters

import (
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	DB *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) usecases.UserRepo {
	return &GormUserRepo{DB: db}
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

func (r *GormUserRepo) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepo) UpdateUser(user *entities.User) error {
	return r.DB.Save(user).Error
}

func (r *GormUserRepo) DeleteUser(userID string) error {
	return r.DB.Delete(&entities.User{}, userID).Error
}
