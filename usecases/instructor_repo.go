package usecases

import (
	"github.com/songwaad/cs-event-backend/entities"
)

type InstructorRepo interface {
	Create(instructor *entities.Instructor) error
	GetByID(id int) (*entities.Instructor, error)
	GetAll() ([]entities.Instructor, error)
	Update(instructor *entities.Instructor) error
	Delete(id int) error
}
