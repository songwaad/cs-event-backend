package adapters

import (
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
	"gorm.io/gorm"
)

type GormInstructorRepo struct {
	DB *gorm.DB
}

func NewGormInstructorRepo(DB *gorm.DB) usecases.InstructorRepo {
	return &GormInstructorRepo{DB: DB}
}

func (r *GormInstructorRepo) Create(instructor *entities.Instructor) error {
	return r.DB.Create(instructor).Error
}

func (r *GormInstructorRepo) GetByID(id int) (*entities.Instructor, error) {
	var instructor entities.Instructor
	result := r.DB.First(&instructor, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &instructor, nil
}

func (r *GormInstructorRepo) GetAll() ([]entities.Instructor, error) {
	var instructors []entities.Instructor
	result := r.DB.Find(&instructors)
	if result.Error != nil {
		return nil, result.Error
	}
	return instructors, nil
}

func (r *GormInstructorRepo) Update(instructor *entities.Instructor) error {
	return r.DB.Save(instructor).Error
}

func (r *GormInstructorRepo) Delete(id int) error {
	return r.DB.Delete(&entities.Instructor{}, id).Error
}
