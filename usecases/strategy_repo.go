package usecases

import "github.com/songwaad/cs-event-backend/entities"

type StrategyRepo interface {
	GetAllStrategy() ([]entities.Strategy, error)
	GetStrategyByID(strategyID uint) (*entities.Strategy, error)
	GetAllEventeStrategy() ([]entities.EventStrategy, error)
	GetEventeStrategyByID(eventStrategyID uint) (*entities.EventStrategy, error)
}
