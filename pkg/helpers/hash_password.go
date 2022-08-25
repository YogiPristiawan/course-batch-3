package helpers

import (
	"course/domain"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hashPassword string, password string) (err domain.HttpError) {
	rawError := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if rawError != nil {
		err = domain.NewBadRequestError(fmt.Errorf("password salah"))
	}
	return
}
