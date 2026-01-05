package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func AccessToken(id string, expiredTime int64) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": expiredTime,
	})

	accessString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))

	if err != nil {
		return accessString, err
	}

	return accessString, nil
}
