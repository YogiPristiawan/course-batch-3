package rest

import (
	"course/domain"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type exerciseValidator struct {
	validator *validator.Validate
}

func NewExerciseValidator(
	validator *validator.Validate,
) *exerciseValidator {
	return &exerciseValidator{
		validator: validator,
	}
}

func (e *exerciseValidator) ValidateCreateExercisePayload(payload *domain.ExerciseCreateRequest) (err error) {
	err = e.validator.Struct(payload)

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("%s harus diisi", err.Field()))
			case "max":
				return errors.New(fmt.Sprintf("%s tidak boleh lebih dari %s karakter", err.Field(), err.Param()))
			}
		}
	}
	return
}
