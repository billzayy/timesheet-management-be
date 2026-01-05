package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func RefreshToken(id string) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 168).Unix(), // Exp 7 days
	})

	refreshString, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_KEY")))

	if err != nil {
		return refreshString, err
	}

	return refreshString, nil
}
