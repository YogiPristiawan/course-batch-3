package repositories

import (
	"course/domain"

	"gorm.io/gorm"
)

type exerciseRepository struct {
	db *gorm.DB
}

func NewExerciseRepository(
	db *gorm.DB,
) *exerciseRepository {
	return &exerciseRepository{
		db: db,
	}
}

func (e *exerciseRepository) Create(exercise *domain.ExerciseModel) (err error) {
	err = e.db.Create(exercise).Error
	return
}
