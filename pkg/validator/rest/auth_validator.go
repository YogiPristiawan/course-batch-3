package rest

import (
	"course/domain"
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

func (a *authValidator) ValidateRegisterRequest(payload *domain.AuthRegisterRequest) (err domain.HttpError) {
	vError := a.validator.Struct(payload)

	if castedObject, ok := vError.(validator.ValidationErrors); ok {
		for _, vError := range castedObject {
			switch vError.Tag() {
			case "required":
				return domain.NewBadRequestError(fmt.Errorf("%s harus diisi", vError.Field()))
			case "email":
				return domain.NewBadRequestError(fmt.Errorf("%s harus berupa valid email", vError.Field()))
			case "gt":
				return domain.NewBadRequestError(fmt.Errorf("%s harus lebih dari %s karakter", vError.Field(), vError.Param()))
			case "numeric":
				return domain.NewBadRequestError(fmt.Errorf("%s harus berupa angka", vError.Field()))
			}
		}
	}
	return
}

func (a *authValidator) ValidateLoginRequest(payload *domain.AuthLoginRequest) (err domain.HttpError) {
	vError := a.validator.Struct(payload)

	if castedObject, ok := vError.(validator.ValidationErrors); ok {
		for _, vError := range castedObject {
			switch vError.Tag() {
			case "required":
				return domain.NewBadRequestError(fmt.Errorf("%s harus diisi", vError.Field()))
			case "email":
				return domain.NewBadRequestError(fmt.Errorf("%s harus berupa email yang valid", vError.Field()))
			}
		}
	}
	return
}
