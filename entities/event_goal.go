package entities

type EventGoal struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Goal string `json:"goal"`
}
