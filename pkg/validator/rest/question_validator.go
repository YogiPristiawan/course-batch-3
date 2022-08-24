package rest

import (
	"course/domain"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type questionValidator struct {
	validator *validator.Validate
}

func NewQuestionValidator(
	validator *validator.Validate,
) domain.QuestionValidator {
	return &questionValidator{
		validator: validator,
	}
}

func (q *questionValidator) ValidateCreateQuestionPayload(payload *domain.ExerciseQuestionCreateRequest) domain.HttpError {
	err := q.validator.Struct(payload)

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				return domain.NewBadRequestError(fmt.Errorf("%s harus diisi", err.Field()))
			case "eq":
				return domain.NewBadRequestError(fmt.Errorf("%s hanya boleh a,b,c atau d", err.Field()))
			}
		}
	}
	return nil
}
