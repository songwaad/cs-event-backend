package usecases

import "github.com/songwaad/cs-event-backend/entities"

type SpeakerUseCase interface {
	CreateSpeaker(speaker *entities.Speaker) error
	GetSpeakerByID(speakerID uint) (*entities.Speaker, error)
	GetAllSpeakers() ([]entities.Speaker, error)
	UpdateSpeaker(speaker *entities.Speaker) error
	DeleteSpeaker(speakerID uint) error
}

type SpeakerService struct {
	repo SpeakerRepo
}

func NewSpeakerService(repo SpeakerRepo) *SpeakerService {
	return &SpeakerService{repo: repo}
}

func (s *SpeakerService) CreateSpeaker(speaker *entities.Speaker) error {
	return s.repo.Create(speaker)
}

func (s *SpeakerService) GetSpeakerByID(speakerID uint) (*entities.Speaker, error) {
	return s.repo.GetByID(speakerID)
}

func (s *SpeakerService) GetAllSpeakers() ([]entities.Speaker, error) {
	return s.repo.GetAll()
}

func (s *SpeakerService) UpdateSpeaker(speaker *entities.Speaker) error {
	return s.repo.Update(speaker)
}

func (s *SpeakerService) DeleteSpeaker(speakerID uint) error {
	return s.repo.Delete(speakerID)
}
