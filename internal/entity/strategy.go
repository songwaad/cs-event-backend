package entity

type Strategy struct {
	StrategyID uint `json:"id" gorm:"primaryKey"`
	Strategy   string
}
