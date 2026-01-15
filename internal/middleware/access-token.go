package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func AccessToken(id string) (string, int64, error) {
	expiredTime := time.Now().Add(time.Hour * 24).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": expiredTime,
	})

	accessString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))

	if err != nil {
		return accessString, 0, err
	}

	return accessString, expiredTime, nil
}
