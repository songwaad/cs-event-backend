package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	DB *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) UserRepo {
	return &GormUserRepo{DB: db}
}

func (r *GormUserRepo) CreateUser(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepo) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Preload("UserRole").Preload("UserStatus").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepo) GetUserByID(userID string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Preload("UserRole").Preload("UserStatus").Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepo) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.DB.Preload("UserRole").Preload("UserStatus").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepo) UpdateUser(user *entity.User) error {
	return r.DB.Model(&entity.User{}).Where("user_id = ?", user.UserID).Save(*user).Error
}

func (r *GormUserRepo) UpdateUserRole(userID string, userRoleID uint) error {
	return r.DB.Model(&entity.User{}).Where("user_id = ?", userID).Updates(map[string]interface{}{
		"user_role_id": userRoleID,
	}).Error
}

func (r *GormUserRepo) UpdateUserStatus(userID string, userStatusID uint) error {
	return r.DB.Model(&entity.User{}).Where("user_id = ?", userID).Updates(map[string]interface{}{
		"user_status_id": userStatusID,
	}).Error
}

func (r *GormUserRepo) DeleteUser(userID string) error {
	return r.DB.Delete(&entity.User{}, userID).Error
}
