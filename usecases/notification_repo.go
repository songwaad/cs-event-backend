package usecases

import "github.com/songwaad/cs-event-backend/entities"

type NotificationRepo interface {
	GetByUserID(userID string) ([]entities.Notification, error)
	InActive() error
	Delete(id uint) error
}
