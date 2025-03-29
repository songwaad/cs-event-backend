package entities

type Strategy struct {
	StrategyID uint `json:"id" gorm:"primaryKey"`
	Strategy   string
}
