package repositories

import (
	"course/domain"
	"course/pkg/helpers"

	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(
	db *gorm.DB,
) domain.QuestionRepository {
	return &questionRepository{
		db: db,
	}
}

func (q *questionRepository) FindByExerciseId(exerciseId int) (questions []*domain.QuestionModel, err domain.HttpError) {
	dbError := q.db.Where("exercise_id = ?", exerciseId).Find(&questions).Error
	err = helpers.CastDatabaseError(dbError, true)
	return
}
