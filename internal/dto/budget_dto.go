package dto

import "github.com/songwaad/cs-event-backend/internal/entity"

type BudgetDTO struct {
	BudgetID    uint `json:"id" gorm:"primaryKey"`
	EventID     uint
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func ToBudgetResponseDTO(budget *entity.Budget) BudgetDTO {
	return BudgetDTO{
		BudgetID:    budget.BudgetID,
		EventID:     budget.EventID,
		Amount:      budget.Amount,
		Description: budget.Description,
	}
}
