package entities

type Budget struct {
	BudgetID    uint `json:"budget_id" gorm:"primaryKey"`
	EventID     uint `json:"event_id"`
	Event       Event
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
