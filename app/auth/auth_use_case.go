package auth

import (
	"course/domain"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userRepository         domain.UserRepository
	generateAccessToken    func(int) (string, domain.HttpError)
	compareHashAndPassword func(string, string) domain.HttpError
}

func NewAuthUseCase(
	userRepository domain.UserRepository,
	generateAccessToken func(int) (string, domain.HttpError),
	compareHashAndPassword func(string, string) domain.HttpError,
) domain.AuthUseCase {
	return &authUseCase{
		userRepository:         userRepository,
		generateAccessToken:    generateAccessToken,
		compareHashAndPassword: compareHashAndPassword,
	}
}

func (a *authUseCase) Register(in *domain.AuthRegisterRequest) (out domain.AuthRegisterResponse) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	// verify if email available
	err = a.userRepository.VerifyAvailableEmail(in.Email)
	if domain.HandleHttpError(err, &out.CommonResult) {
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
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	// generate access token
	accessToken, err := a.generateAccessToken(user.ID)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	out.Token = accessToken
	return
}

func (a *authUseCase) Login(in *domain.AuthLoginRequest) (out domain.AuthLoginResponse) {
	// get user by email
	user, err := a.userRepository.GetUserByEmail(in.Email, "id", "email", "password")
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	// compare password
	err = a.compareHashAndPassword(user.Password, in.Password)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	// generate token
	token, rawError := a.generateAccessToken(user.ID)
	if domain.HandleHttpError(rawError, &out.CommonResult) {
		return
	}

	out.Token = token
	return
}
