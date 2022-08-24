package repositories

import (
	"course/domain"
	"course/pkg/helpers"
	"fmt"

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
	err = helpers.CastDatabaseError(dbError, false)
	return
}

func (q *questionRepository) Create(question *domain.QuestionModel) (err domain.HttpError) {
	dbError := q.db.Create(question).Error
	err = helpers.CastDatabaseError(dbError, true)
	return
}

func (q *questionRepository) VerifyExerciseAndQuestionId(exerciseId int, questionId int) (err domain.HttpError) {
	dbError := q.db.Where("id = ?", questionId).Where("exercise_id = ?", exerciseId).First(&domain.QuestionModel{}).Error
	err = helpers.CastDatabaseError(dbError, true)
	return
}

func (q *questionRepository) CreateQuestionAnswer(answer *domain.AnswerModel) (err domain.HttpError) {
	dbError := q.db.Create(answer).Error
	err = helpers.CastDatabaseError(dbError, true)
	return
}

func (q *questionRepository) VerifyExistsAnswer(userId int, exerciseId int, questionId int) (err domain.HttpError) {
	var count int64
	dbError := q.db.
		Model(&domain.AnswerModel{}).
		Where("user_id = ?", userId).
		Where("exercise_id = ?", exerciseId).
		Where("question_id = ?", questionId).Count(&count).Error
	err = helpers.CastDatabaseError(dbError, true)

	if count > 0 {
		err = domain.NewBadRequestError(fmt.Errorf("jawaban sudah di input"))
	}
	return
}
