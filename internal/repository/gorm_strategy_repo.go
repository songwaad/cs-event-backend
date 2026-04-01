package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"gorm.io/gorm"
)

type GormStrategyRepo struct {
	DB *gorm.DB
}

func NewGormStrategyRepo(DB *gorm.DB) StrategyRepo {
	return &GormStrategyRepo{DB: DB}
}

func (r *GormStrategyRepo) GetAllStrategy() ([]entity.Strategy, error) {
	var strategies []entity.Strategy
	result := r.DB.Find(&strategies)
	if result.Error != nil {
		return nil, result.Error
	}
	return strategies, nil
}

func (r *GormStrategyRepo) GetStrategyByID(strategyID uint) (*entity.Strategy, error) {
	var strategy entity.Strategy
	result := r.DB.First(&strategy, strategyID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &strategy, nil
}

func (r *GormStrategyRepo) GetAllEventStrategy() ([]entity.EventStrategy, error) {
	var eventStrategies []entity.EventStrategy
	result := r.DB.Preload("Strategy").Find(&eventStrategies)
	if result.Error != nil {
		return nil, result.Error
	}
	return eventStrategies, nil
}

func (r *GormStrategyRepo) GetEventStrategyByID(eventStrategyID uint) (*entity.EventStrategy, error) {
	var eventStrategy entity.EventStrategy
	result := r.DB.Preload("Strategy").First(&eventStrategy, eventStrategyID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &eventStrategy, nil
}
