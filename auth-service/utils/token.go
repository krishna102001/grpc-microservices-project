package utils

import (
	"os"
	"time"

	logger "github.com/charmbracelet/log"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	var jwtKey = os.Getenv("JWT_KEY")
	if jwtKey == "" {
		logger.Error("Jwt key is missing")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claim := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(jwtKey)
}
