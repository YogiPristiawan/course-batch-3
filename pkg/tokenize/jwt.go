package tokenize

import (
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
