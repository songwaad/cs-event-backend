package adapters

import (
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
	"gorm.io/gorm"
)

type GormSpeakerRepo struct {
	DB *gorm.DB
}

func NewGormSpeakerRepo(DB *gorm.DB) usecases.SpeakerRepo {
	return &GormSpeakerRepo{DB: DB}
}

func (r *GormSpeakerRepo) Create(speaker *entities.Speaker) error {
	return r.DB.Create(speaker).Error
}

func (r *GormSpeakerRepo) GetByID(id int) (*entities.Speaker, error) {
	var speaker entities.Speaker
	result := r.DB.First(&speaker, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &speaker, nil
}

func (r *GormSpeakerRepo) GetAll() ([]entities.Speaker, error) {
	var speakers []entities.Speaker
	result := r.DB.Find(&speakers)
	if result.Error != nil {
		return nil, result.Error
	}
	return speakers, nil
}

func (r *GormSpeakerRepo) Update(speaker *entities.Speaker) error {
	return r.DB.Save(speaker).Error
}

func (r *GormSpeakerRepo) Delete(id int) error {
	return r.DB.Delete(&entities.Speaker{}, id).Error
}
