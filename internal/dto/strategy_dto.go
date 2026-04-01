package dto

import "github.com/songwaad/cs-event-backend/internal/entity"

type EventStrategyResponseDTO struct {
	EventStrategyID uint `json:"event_strategy_id" gorm:"primaryKey"`
	Strategy        string
	Goal            string `json:"goal"`
	Tactic          string `json:"tactic"`
}

type StrategyResponseDTO struct {
	StrategyID uint   `json:"strategy_id"`
	Strategy   string `json:"strategy"`
}

func ToEventStrategyResponseDTO(eventStrategy entity.EventStrategy) EventStrategyResponseDTO {
	return EventStrategyResponseDTO{
		EventStrategyID: eventStrategy.EventStrategyID,
		Strategy:        eventStrategy.Strategy.Strategy,
		Goal:            eventStrategy.Goal,
		Tactic:          eventStrategy.Tactic,
	}
}

func ToStrategyResponseDTO(strategy entity.Strategy) StrategyResponseDTO {
	return StrategyResponseDTO{
		StrategyID: strategy.StrategyID,
		Strategy:   strategy.Strategy,
	}
}
