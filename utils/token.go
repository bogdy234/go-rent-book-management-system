package utils

import (
	"os"

	"github.com/golang-jwt/jwt"
)

// Encode passed in data as JWT
func EncodeDataJWT(j jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}

	return tokenString
}
