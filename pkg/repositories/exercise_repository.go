package repositories

import (
	"course/domain"
	"course/pkg/helpers"

	"gorm.io/gorm"
)

type exerciseRepository struct {
	db *gorm.DB
}

func NewExerciseRepository(
	db *gorm.DB,
) domain.ExerciseRepository {
	return &exerciseRepository{
		db: db,
	}
}

func (e *exerciseRepository) Create(exercise *domain.ExerciseModel) (err error) {
	err = e.db.Create(exercise).Error
	return
}

func (e *exerciseRepository) GetById(id int) (result *domain.ExerciseModel, err domain.HttpError) {
	dbError := e.db.Where("id = ?", id).First(&result).Error
	err = helpers.CastDatabaseError(dbError, true)
	return
}

func (e *exerciseRepository) FindUserQuestionAnswer(exerciseId int, userId int) (results []map[string]interface{}, err domain.HttpError) {
	dbError := e.db.Raw(`
		SELECT
			questions.correct_answer,
			questions.score,
			answers.answer AS user_answer
		FROM
			answers
			INNER JOIN questions ON questions.id = answers.question_id AND questions.exercise_id = ?
		WHERE
			answers.user_id = ?
	`, exerciseId, userId).Find(&results).Error
	err = helpers.CastDatabaseError(dbError, false)
	return
}
