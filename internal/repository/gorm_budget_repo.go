package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"gorm.io/gorm"
)

type GormBudgetRepo struct {
	DB *gorm.DB
}

func NewGormBudgetRepo(DB *gorm.DB) BudgetRepo {
	return &GormBudgetRepo{DB: DB}
}

func (r *GormBudgetRepo) Create(budget *entity.Budget) error {
	return r.DB.Create(budget).Error
}

func (r *GormBudgetRepo) GetByID(id uint) (*entity.Budget, error) {
	var budget entity.Budget
	result := r.DB.First(&budget, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &budget, nil
}

func (r *GormBudgetRepo) GetAll() ([]entity.Budget, error) {
	var budgets []entity.Budget
	result := r.DB.Find(&budgets)
	if result.Error != nil {
		return nil, result.Error
	}
	return budgets, nil
}

func (r *GormBudgetRepo) Update(budget *entity.Budget) error {
	return r.DB.Save(budget).Error
}

func (r *GormBudgetRepo) Delete(id uint) error {
	return r.DB.Delete(&entity.Budget{}, id).Error
}
