package rest

import (
	"course/domain"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type exerciseValidator struct {
	validator *validator.Validate
}

func NewExerciseValidator(
	validator *validator.Validate,
) domain.ExerciseValidator {
	return &exerciseValidator{
		validator: validator,
	}
}

func (e *exerciseValidator) ValidateCreateExercisePayload(payload *domain.ExerciseCreateRequest) (err domain.HttpError) {
	vError := e.validator.Struct(payload)

	if castedObject, ok := vError.(validator.ValidationErrors); ok {
		for _, vError := range castedObject {
			switch vError.Tag() {
			case "required":
				return domain.NewBadRequestError(fmt.Errorf("%s harus diisi", vError.Field()))
			case "max":
				return domain.NewBadRequestError(fmt.Errorf("%s tidak boleh lebih dari %s karakter", vError.Field(), vError.Param()))
			}
		}
	}
	return
}
