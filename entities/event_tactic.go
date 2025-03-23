package entities

type EventTactic struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Tactic string `json:"tactic"`
}
