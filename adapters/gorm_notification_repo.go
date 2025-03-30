package adapters

import (
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
	"gorm.io/gorm"
)

type GormNotificationRepo struct {
	DB *gorm.DB
}

func NewGormNotificationRepo(DB *gorm.DB) usecases.NotificationRepo {
	return &GormNotificationRepo{DB: DB}
}

func (r *GormNotificationRepo) GetByUserID(userID string) ([]entities.Notification, error) {
	var notifications []entities.Notification
	result := r.DB.Where("userID = ?", userID).Preload("Event").Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifications, nil
}

func (r *GormNotificationRepo) CreateNotify(notify *entities.Notification) error {
	return r.DB.Create(notify).Error
}

func (r *GormNotificationRepo) InActive(id uint) error {
	result := r.DB.Model(&entities.Notification{}).
		Where("id = ?", id).
		Update("active", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormNotificationRepo) Delete(id uint) error {
	return r.DB.Delete(&entities.Notification{}, id).Error
}
