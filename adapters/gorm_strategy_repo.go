package adapters

import (
	"github.com/songwaad/cs-event-backend/entities"
	"gorm.io/gorm"
)

type GormStrategyRepo struct {
	DB *gorm.DB
}

func NewGormStrategyRepo(DB *gorm.DB) *GormStrategyRepo {
	return &GormStrategyRepo{DB: DB}
}

func (r *GormStrategyRepo) GetAllStrategy() ([]entities.Strategy, error) {
	var stragies []entities.Strategy
	result := r.DB.Find(&stragies)
	if result.Error != nil {
		return nil, result.Error
	}
	return stragies, nil
}

func (r *GormStrategyRepo) GetStrategyByID(strategyID uint) (*entities.Strategy, error) {
	var strategy entities.Strategy
	result := r.DB.First(&strategy, strategyID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &strategy, nil
}

func (r *GormStrategyRepo) GetAllEventeStrategy() ([]entities.EventStrategy, error) {
	var eventStragies []entities.EventStrategy
	result := r.DB.Preload("Strategy").Find(&eventStragies)
	if result.Error != nil {
		return nil, result.Error
	}
	return eventStragies, nil
}

func (r *GormStrategyRepo) GetEventeStrategyByID(eventStrategyID uint) (*entities.EventStrategy, error) {
	var eventStrategy entities.EventStrategy
	result := r.DB.Preload("Strategy").First(&eventStrategy, eventStrategyID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &eventStrategy, nil
}
