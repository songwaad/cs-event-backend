package entities

type EventStrategy struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	StrategyID uint
	Strategy   Strategy
	Goal       string `json:"goal"`
	Tactic     string `json:"tactic"`
}
