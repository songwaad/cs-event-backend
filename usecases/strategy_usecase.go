package usecases

import "github.com/songwaad/cs-event-backend/entities"

type StrategyUsecases interface {
	GetAllStrategy() ([]entities.Strategy, error)
	GetStrategyByID(strategyID uint) (*entities.Strategy, error)
	GetAllEventeStrategy() ([]entities.EventStrategy, error)
	GetEventeStrategyByID(eventStrategyID uint) (*entities.EventStrategy, error)
}

type StrategyService struct {
	repo StrategyRepo
}

func NewStrategyService(repo StrategyRepo) *StrategyService {
	return &StrategyService{repo: repo}
}

func (s *StrategyService) GetAllStrategy() ([]entities.Strategy, error) {
	return s.repo.GetAllStrategy()
}

func (s *StrategyService) GetStrategyByID(strategyID uint) (*entities.Strategy, error) {
	return s.repo.GetStrategyByID(strategyID)
}

func (s *StrategyService) GetAllEventeStrategy() ([]entities.EventStrategy, error) {
	return s.repo.GetAllEventeStrategy()
}

func (s *StrategyService) GetEventeStrategyByID(eventStrategyID uint) (*entities.EventStrategy, error) {
	return s.repo.GetEventeStrategyByID(eventStrategyID)
}
