package usecases

import "github.com/songwaad/cs-event-backend/entities"

type NotificationRepo interface {
	CreateNotify(notify *entities.Notification) error
	GetByUserID(userID string) ([]entities.Notification, error)
	InActive(id uint) error
	Delete(id uint) error
}
