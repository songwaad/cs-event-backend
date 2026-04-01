package usecase

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"github.com/songwaad/cs-event-backend/internal/repository"
)

type SpeakerUseCase interface {
	CreateSpeaker(speaker *entity.Speaker) error
	GetSpeakerByID(speakerID uint) (*entity.Speaker, error)
	GetAllSpeakers() ([]entity.Speaker, error)
	UpdateSpeaker(speaker *entity.Speaker) error
	DeleteSpeaker(speakerID uint) error
}

type SpeakerService struct {
	repo repository.SpeakerRepo
}

func NewSpeakerService(repo repository.SpeakerRepo) *SpeakerService {
	return &SpeakerService{repo: repo}
}

func (s *SpeakerService) CreateSpeaker(speaker *entity.Speaker) error {
	return s.repo.Create(speaker)
}

func (s *SpeakerService) GetSpeakerByID(speakerID uint) (*entity.Speaker, error) {
	return s.repo.GetByID(speakerID)
}

func (s *SpeakerService) GetAllSpeakers() ([]entity.Speaker, error) {
	return s.repo.GetAll()
}

func (s *SpeakerService) UpdateSpeaker(speaker *entity.Speaker) error {
	return s.repo.Update(speaker)
}

func (s *SpeakerService) DeleteSpeaker(speakerID uint) error {
	return s.repo.Delete(speakerID)
}
