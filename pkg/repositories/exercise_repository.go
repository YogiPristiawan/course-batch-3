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
	err = helpers.CastDatabaseError(dbError, false)
	return
}

func (e *exerciseRepository) GetByIdWithQuestions(id int) (results []map[string]interface{}, err error) {
	err = e.db.Raw(`
		SELECT
			exercises.id,
			exercises.title,
			exercises.description,
			questions.id AS question_id,
			questions.body AS question_body,
			questions.option_a AS question_option_a,
			questions.option_b AS question_option_b,
			questions.option_c AS question_option_c,
			questions.option_d AS question_option_d,
			questions.score AS question_score,
			questions.created_at AS question_created_at,
			questions.updated_at AS question_updated_at
		FROM
			exercises
			JOIN questions ON questions.exercise_id = exercises.id AND exercises.id = ?
		ORDER BY
			questions.id DESC
	`, id).Scan(&results).Error
	return
}
