package jwt

import (
	"time"

	golangJwt "github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userID string, role string, secretKey string) (string, error) {
	token := golangJwt.New(golangJwt.SigningMethodHS256)
	claims := token.Claims.(golangJwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(secretKey))
}
