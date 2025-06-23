package dto

import "github.com/songwaad/cs-event-backend/entities"

type SpeakerDTO struct {
	SpeakerID   uint   `json:"speaker_id" gorm:"primaryKey"`
	FirstName   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Description string `json:"description"`
	ImageUrl    string
}

func ToSpeakerResponseDTO(entity entities.Speaker) SpeakerDTO {
	return SpeakerDTO{
		SpeakerID:   entity.SpeakerID,
		FirstName:   entity.FirstName,
		Lastname:    entity.Lastname,
		Description: entity.Description,
		ImageUrl:    entity.ImageUrl,
	}
}
