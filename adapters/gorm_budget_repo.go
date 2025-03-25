package adapters

import (
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
	"gorm.io/gorm"
)

type GormBudgetRepo struct {
	DB *gorm.DB
}

func NewGormBudgetRepo(DB *gorm.DB) usecases.BudgetRepo {
	return &GormBudgetRepo{DB: DB}
}

func (r *GormBudgetRepo) Create(budget *entities.Budget) error {
	return r.DB.Create(budget).Error
}

func (r *GormBudgetRepo) GetByID(id uint) (*entities.Budget, error) {
	var budget entities.Budget
	result := r.DB.First(&budget, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &budget, nil
}

func (r *GormBudgetRepo) GetAll() ([]entities.Budget, error) {
	var budgets []entities.Budget
	result := r.DB.Find(&budgets)
	if result.Error != nil {
		return nil, result.Error
	}
	return budgets, nil
}

func (r *GormBudgetRepo) Update(budget *entities.Budget) error {
	return r.DB.Save(budget).Error
}

func (r *GormBudgetRepo) Delete(id uint) error {
	return r.DB.Delete(&entities.Budget{}, id).Error
}
