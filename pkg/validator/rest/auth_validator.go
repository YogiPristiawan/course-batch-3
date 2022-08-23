package rest

import (
	"course/domain"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type authValidator struct {
	validator *validator.Validate
}

func NewAuthValidator(validator *validator.Validate) domain.AuthValidator {
	return &authValidator{
		validator: validator,
	}
}

func (a *authValidator) ValidateRegisterRequest(payload *domain.AuthRegisterRequest) (err error) {
	err = a.validator.Struct(payload)

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("%s harus di isi", err.Field()))
			case "email":
				return errors.New(fmt.Sprintf("%s harus berupa valid email", err.Field()))
			case "gt":
				return errors.New(fmt.Sprintf("%s harus lebih dari %s karakter", err.Field(), err.Param()))
			case "numeric":
				return errors.New(fmt.Sprintf("%s harus berupa angka", err.Field()))
			}
		}
	}
	return
}
