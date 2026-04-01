package repository

import (
	"github.com/songwaad/cs-event-backend/internal/entity"
	"gorm.io/gorm"
)

type GormSpeakerRepo struct {
	DB *gorm.DB
}

func NewGormSpeakerRepo(DB *gorm.DB) SpeakerRepo {
	return &GormSpeakerRepo{DB: DB}
}

func (r *GormSpeakerRepo) Create(speaker *entity.Speaker) error {
	return r.DB.Create(speaker).Error
}

func (r *GormSpeakerRepo) GetByID(speakerID uint) (*entity.Speaker, error) {
	var speaker entity.Speaker
	result := r.DB.First(&speaker, speakerID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &speaker, nil
}

func (r *GormSpeakerRepo) GetAll() ([]entity.Speaker, error) {
	var speakers []entity.Speaker
	result := r.DB.Find(&speakers)
	if result.Error != nil {
		return nil, result.Error
	}
	return speakers, nil
}

func (r *GormSpeakerRepo) Update(speaker *entity.Speaker) error {
	return r.DB.Save(speaker).Error
}

func (r *GormSpeakerRepo) Delete(speakerID uint) error {
	return r.DB.Delete(&entity.Speaker{}, speakerID).Error
}
