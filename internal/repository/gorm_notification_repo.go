package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"gorm.io/gorm"
)

type GormNotificationRepo struct {
	DB *gorm.DB
}

func NewGormNotificationRepo(DB *gorm.DB) NotificationRepo {
	return &GormNotificationRepo{DB: DB}
}

func (r *GormNotificationRepo) GetByUserID(userID string) ([]entity.Notification, error) {
	var notifications []entity.Notification
	result := r.DB.Where("user_id = ?", userID).Preload("Event").Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifications, nil
}

func (r *GormNotificationRepo) CreateNotify(notify *entity.Notification) error {
	return r.DB.Create(notify).Error
}

func (r *GormNotificationRepo) InActive(id uint) error {
	result := r.DB.Model(&entity.Notification{}).
		Where("notify_id = ?", id).
		Update("active", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormNotificationRepo) Delete(id uint) error {
	return r.DB.Delete(&entity.Notification{}, id).Error
}
