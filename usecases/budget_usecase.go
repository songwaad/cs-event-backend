package usecases

import (
	"github.com/songwaad/cs-event-backend/entities"
)

type BudgetUseCase interface {
	CreateBudget(budget *entities.Budget) error
	GetBudgetByID(id uint) (*entities.Budget, error)
	GetALLBudgets() ([]entities.Budget, error)
	UpdateBudget(budget *entities.Budget) error
	DeleteBudget(id uint) error
}

type BudgetService struct {
	repo BudgetRepo
}

func NewBudgetService(repo BudgetRepo) *BudgetService {
	return &BudgetService{repo: repo}
}

func (s *BudgetService) CreateBudget(budget *entities.Budget) error {
	return s.repo.Create(budget)
}

func (s *BudgetService) GetBudgetByID(id uint) (*entities.Budget, error) {
	return s.repo.GetByID(id)
}

func (s *BudgetService) GetALLBudgets() ([]entities.Budget, error) {
	return s.repo.GetAll()
}

func (s *BudgetService) UpdateBudget(budget *entities.Budget) error {
	return s.repo.Update(budget)
}

func (s *BudgetService) DeleteBudget(id uint) error {
	return s.repo.Delete(id)
}
