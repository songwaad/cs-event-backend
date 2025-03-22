package usecases

import "github.com/songwaad/cs-event-backend/entities"

type InstructorUseCase interface {
	CreateInstructor(instructor *entities.Instructor) error
	GetInstructorByID(id int) (*entities.Instructor, error)
	GetAllInstructors() ([]entities.Instructor, error)
	UpdateInstructor(instructor *entities.Instructor) error
	DeleteInstructor(id int) error
}

type InstructorService struct {
	repo InstructorRepo
}

func NewInstructorService(repo InstructorRepo) *InstructorService {
	return &InstructorService{repo: repo}
}

func (s *InstructorService) CreateInstructor(instructor *entities.Instructor) error {
	return s.repo.Create(instructor)
}

func (s *InstructorService) GetInstructorByID(id int) (*entities.Instructor, error) {
	return s.repo.GetByID(id)
}

func (s *InstructorService) GetAllInstructors() ([]entities.Instructor, error) {
	return s.repo.GetAll()
}

func (s *InstructorService) UpdateInstructor(instructor *entities.Instructor) error {
	return s.repo.Update(instructor)
}

func (s *InstructorService) DeleteInstructor(id int) error {
	return s.repo.Delete(id)
}
