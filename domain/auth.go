package domain

type AuthUseCase interface {
	Register(*AuthRegisterRequest) AuthRegisterResponse
	Login(*AuthLoginRequest) AuthLoginResponse
}

type AuthValidator interface {
	ValidateRegisterRequest(*AuthRegisterRequest) HttpError
	ValidateLoginRequest(*AuthLoginRequest) HttpError
}

type AuthRegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gt=6"`
	NoHp     string `json:"no_hp" validate:"required,numeric"`
}

type AuthRegisterResponse struct {
	CommonResult
	Token string `json:"token"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginResponse struct {
	CommonResult
	Token string `json:"token"`
}
