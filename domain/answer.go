package domain

import "time"

type AnswerModel struct {
	ID         int `gorm:"primaryKey"`
	ExerciseId int
	QuestionId int
	UserId     int
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (a *AnswerModel) TableName() string {
	return "answers"
}
