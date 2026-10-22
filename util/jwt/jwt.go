package jwt

import (
	"time"

	"github.com/Mikhail-Tal63/Rooksy/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID any) (string, error) {
	exp := time.Second * time.Duration(config.Load().JwtExp)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(exp).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
