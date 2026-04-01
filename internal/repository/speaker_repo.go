package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
)

type SpeakerRepo interface {
	Create(speaker *entity.Speaker) error
	GetByID(speakerID uint) (*entity.Speaker, error)
	GetAll() ([]entity.Speaker, error)
	Update(speakerID *entity.Speaker) error
	Delete(speakerID uint) error
}
