package repository

import "github.com/songwaad/cs-event-backend/internal/entity"

type BudgetRepo interface {
	Create(*entity.Budget) error
	GetByID(id uint) (*entity.Budget, error)
	GetAll() ([]entity.Budget, error)
	Update(budget *entity.Budget) error
	Delete(id uint) error
}
