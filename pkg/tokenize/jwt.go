package tokenize

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var AccessTokenSecretKey []byte = []byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY"))

func GenerateAccessToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessTokenSecretKey)
}

func DecodeJwt(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("algoritma token tidak sesuai: %s", token.Header["alg"])
		}

		return AccessTokenSecretKey, nil
	})
	if err != nil {
		return map[string]interface{}{}, err
	}

	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return map[string]interface{}{}, errors.New("token tidak valid")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return map[string]interface{}{}, errors.New("token kedaluwarsa atau tidak aktif")
			} else {
				return map[string]interface{}{}, errors.New("error saat decode token")
			}
		}
	}

	return token.Claims.(jwt.MapClaims), nil
}
