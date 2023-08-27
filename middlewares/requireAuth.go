package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"rent-book-management-system/config"
	"rent-book-management-system/models"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	session := sessions.Default(c)

	// Get the token from session
	sessionEncodedToken := session.Get("token")
	if sessionEncodedToken == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not present in the request",
		})
	}

	tokenString := fmt.Sprintf("%v", sessionEncodedToken)

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
		fmt.Println(claims["foo"], claims["nbf"])

		// Check if it's expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user via token sub
		var user models.User
		config.GetDB().First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach it to the request
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		fmt.Println(err)
	}

}
