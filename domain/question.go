package domain

import (
	"time"
)

type QuestionModel struct {
	ID            int `gorm:"primaryKey"`
	ExerciseId    int
	Body          string
	OptionA       string
	OptionB       string
	OptionC       string
	OptionD       string
	CorrectAnswer string
	Score         int
	CreatorId     int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (q *QuestionModel) TableName() string {
	return "questions"
}

type QuestionValidator interface {
	ValidateCreateQuestionPayload(*ExerciseQuestionCreateRequest) HttpError
}

type QuestionRepository interface {
	FindByExerciseId(int) ([]*QuestionModel, HttpError)
	Create(*QuestionModel) HttpError
}
