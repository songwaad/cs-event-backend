package entities

type EventStrategy struct {
	EventStrategyID uint `json:"event_strategy_id" gorm:"primaryKey"`
	StrategyID      uint `json:"stragy_id"`
	Strategy        Strategy
	Goal            string `json:"goal"`
	Tactic          string `json:"tactic"`
}
