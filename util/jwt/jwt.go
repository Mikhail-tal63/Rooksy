package jwt

import (
	"fmt"
	"time"

	"github.com/Mikhail-Tal63/Rooksy/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func GenerateRefreshToken(secret []byte, UserID any) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": UserID,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
func VerifyJWT(secret []byte, tokenString string) (primitive.ObjectID, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return primitive.NilObjectID, err
	}

	if !token.Valid {
		return primitive.NilObjectID, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("invalid token claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("invalid token sub")
	}

	userID, err := primitive.ObjectIDFromHex(sub)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return userID, nil
}
