package repository

import "github.com/songwaad/cs-event-backend/internal/entity"

type NotificationRepo interface {
	CreateNotify(notify *entity.Notification) error
	GetByUserID(userID string) ([]entity.Notification, error)
	InActive(id uint) error
	Delete(id uint) error
}
