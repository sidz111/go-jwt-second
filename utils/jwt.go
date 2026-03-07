package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = []byte("supersecretkey")

func GenerateJWT(username string, userID uint) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"user_id":  userID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "no token string", errors.New("failed to generate token")
	}
	return tokenString, err
}
