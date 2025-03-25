package usecases

import "github.com/songwaad/cs-event-backend/entities"

type BudgetRepo interface {
	Create(*entities.Budget) error
	GetByID(id uint) (*entities.Budget, error)
	GetAll() ([]entities.Budget, error)
	Update(budget *entities.Budget) error
	Delete(id uint) error
}
