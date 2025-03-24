package entities

type Strategy struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	Strategy string
}
