package usecases

import "github.com/songwaad/cs-event-backend/entities"

type NotificationUsecase interface {
	GetNotifyByUserID(userID string) ([]entities.Notification, error)
	InActive(id uint) error
	DeleteNotify(id uint) error
}

type NotificationService struct {
	repo NotificationRepo
}

func NewNotificationService(repo NotificationRepo) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) GetNotifyByUserID(userID string) ([]entities.Notification, error) {
	return s.repo.GetByUserID(userID)
}

func (s *NotificationService) InActive(id uint) error {
	return s.repo.InActive(id)
}

func (s *NotificationService) DeleteNotify(id uint) error {
	return s.repo.Delete(id)
}
