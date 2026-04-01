package repository

import "github.com/songwaad/cs-event-backend/internal/entity"

type StrategyRepo interface {
	GetAllStrategy() ([]entity.Strategy, error)
	GetStrategyByID(strategyID uint) (*entity.Strategy, error)
	GetAllEventStrategy() ([]entity.EventStrategy, error)
	GetEventStrategyByID(eventStrategyID uint) (*entity.EventStrategy, error)
}
