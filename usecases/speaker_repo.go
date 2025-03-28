package usecases

import (
	"github.com/songwaad/cs-event-backend/entities"
)

type SpeakerRepo interface {
	Create(speaker *entities.Speaker) error
	GetByID(id int) (*entities.Speaker, error)
	GetAll() ([]entities.Speaker, error)
	Update(speaker *entities.Speaker) error
	Delete(id int) error
}
