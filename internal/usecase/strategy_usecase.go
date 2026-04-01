package usecase

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"github.com/songwaad/cs-event-backend/internal/repository"
)

type StrategyUsecases interface {
	GetAllStrategy() ([]entity.Strategy, error)
	GetStrategyByID(strategyID uint) (*entity.Strategy, error)
	GetAllEventStrategy() ([]entity.EventStrategy, error)
	GetEventStrategyByID(eventStrategyID uint) (*entity.EventStrategy, error)
}

type StrategyService struct {
	repo repository.StrategyRepo
}

func NewStrategyService(repo repository.StrategyRepo) *StrategyService {
	return &StrategyService{repo: repo}
}

func (s *StrategyService) GetAllStrategy() ([]entity.Strategy, error) {
	return s.repo.GetAllStrategy()
}

func (s *StrategyService) GetStrategyByID(strategyID uint) (*entity.Strategy, error) {
	return s.repo.GetStrategyByID(strategyID)
}

func (s *StrategyService) GetAllEventStrategy() ([]entity.EventStrategy, error) {
	return s.repo.GetAllEventStrategy()
}

func (s *StrategyService) GetEventStrategyByID(eventStrategyID uint) (*entity.EventStrategy, error) {
	return s.repo.GetEventStrategyByID(eventStrategyID)
}
