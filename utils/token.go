package utils

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
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

func BaseDecode(tokenString string) (jwt.MapClaims, bool) {
	// Decode and validate it
	byteArrSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return byteArrSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return nil, false
	}
}

// GetTokenFromSession returns JWT token from session as string and a bool if it exists on the session or not
func GetTokenFromSession(session sessions.Session) (string, bool) {
	// Get the token from session
	sessionEncodedToken := session.Get("token")
	if sessionEncodedToken == nil {
		return "", false
	}

	return fmt.Sprintf("%v", sessionEncodedToken), true
}
