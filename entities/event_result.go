package entities

import "time"

type EventResult struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	DeletedAt         *time.Time `gorm:"index"`
	Product           string
	ProductIndicators string
	ProductTarget     string
	Result            string
	ResultIndicators  string
	ResultTarget      string
}
