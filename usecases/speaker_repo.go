package usecases

import (
	"github.com/songwaad/cs-event-backend/entities"
)

type SpeakerRepo interface {
	Create(speaker *entities.Speaker) error
	GetByID(speakerID uint) (*entities.Speaker, error)
	GetAll() ([]entities.Speaker, error)
	Update(speakerID *entities.Speaker) error
	Delete(speakerID uint) error
}
