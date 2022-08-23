package auth

import (
	"course/domain"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userRepository      domain.UserRepository
	generateAccessToken func(int) (string, error)
}

func NewAuthUseCase(
	userRepository domain.UserRepository,
	generateAccessToken func(int) (string, error),
) domain.AuthUseCase {
	return &authUseCase{
		userRepository:      userRepository,
		generateAccessToken: generateAccessToken,
	}
}

func (a *authUseCase) Register(in *domain.AuthRegisterRequest) (out domain.AuthRegisterResponse) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		out.SetError(500, err.Error())
		return
	}

	// verify if email available
	if a.userRepository.VerifyAvailableEmail(in.Email) > 0 {
		out.SetError(400, "email sudah terdaftar")
		return
	}

	// insert user
	user := &domain.UserModel{
		Name:     in.Name,
		Email:    in.Email,
		Password: string(hashedPassword),
		NoHp:     in.NoHp,
	}

	err = a.userRepository.Create(user)
	if err != nil {
		out.SetError(500, err.Error())
		return
	}

	// generate access token
	accessToken, err := a.generateAccessToken(user.ID)
	if err != nil {
		out.SetError(500, err.Error())
		return
	}

	out.Token = accessToken
	return
}
