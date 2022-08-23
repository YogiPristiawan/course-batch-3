package domain

type AuthUseCase interface {
	Register(*AuthRegisterRequest) AuthRegisterResponse
}

type AuthValidator interface {
	ValidateRegisterRequest(*AuthRegisterRequest) error
}

type AuthRegisterRequest struct {
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,gt=6"`
	NoHp     *string `json:"no_hp" validate:"required,numeric"`
}

type AuthRegisterResponse struct {
	CommonResult
	Token string `json:"token"`
}
