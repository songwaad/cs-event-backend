package dto

import "github.com/songwaad/cs-event-backend/entities"

type NotifyResponseDTO struct {
	NotifyID uint   `json:"id" gorm:"primaryKey"`
	Active   bool   `json:"active"`
	UserID   string `json:"user_id"`
	EventID  uint   `json:"event_id"`
}

func ToNotifyResponseDTP(notify *entities.Notification) NotifyResponseDTO {
	return NotifyResponseDTO{
		NotifyID: notify.NotifyID,
		Active:   notify.Active,
		UserID:   notify.UserID,
		EventID:  notify.EventID,
	}
}
