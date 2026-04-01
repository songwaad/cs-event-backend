package usecase

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"github.com/songwaad/cs-event-backend/internal/repository"
)

type BudgetUseCase interface {
	CreateBudget(budget *entity.Budget) error
	GetBudgetByID(id uint) (*entity.Budget, error)
	GetALLBudgets() ([]entity.Budget, error)
	UpdateBudget(budget *entity.Budget) error
	DeleteBudget(id uint) error
}

type BudgetService struct {
	repo repository.BudgetRepo
}

func NewBudgetService(repo repository.BudgetRepo) *BudgetService {
	return &BudgetService{repo: repo}
}

func (s *BudgetService) CreateBudget(budget *entity.Budget) error {
	return s.repo.Create(budget)
}

func (s *BudgetService) GetBudgetByID(id uint) (*entity.Budget, error) {
	return s.repo.GetByID(id)
}

func (s *BudgetService) GetALLBudgets() ([]entity.Budget, error) {
	return s.repo.GetAll()
}

func (s *BudgetService) UpdateBudget(budget *entity.Budget) error {
	return s.repo.Update(budget)
}

func (s *BudgetService) DeleteBudget(id uint) error {
	return s.repo.Delete(id)
}
