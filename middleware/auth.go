package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(userId int, username, secretKey string) (string, error) {
	claims := jwtCustomClaims{
		ID:       uint(userId),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
