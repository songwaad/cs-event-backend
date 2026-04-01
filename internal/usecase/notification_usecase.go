package usecase

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"github.com/songwaad/cs-event-backend/internal/repository"
)

type NotificationUsecase interface {
	CreateNotify(notify *entity.Notification) error
	GetNotifyByUserID(userID string) ([]entity.Notification, error)
	InActive(id uint) error
	DeleteNotify(id uint) error
}

type NotificationService struct {
	repo repository.NotificationRepo
}

func NewNotificationService(repo repository.NotificationRepo) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) CreateNotify(notify *entity.Notification) error {
	return s.repo.CreateNotify(notify)
}

func (s *NotificationService) GetNotifyByUserID(userID string) ([]entity.Notification, error) {
	return s.repo.GetByUserID(userID)
}

func (s *NotificationService) InActive(id uint) error {
	return s.repo.InActive(id)
}

func (s *NotificationService) DeleteNotify(id uint) error {
	return s.repo.Delete(id)
}
