package middlewares

import (
	"net/http"
	"rent-book-management-system/config"
	"rent-book-management-system/models"
	"rent-book-management-system/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	session := sessions.Default(c)

	token, ok := utils.GetTokenFromSession(session)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not present in the request",
		})
	}

	claims, _ := utils.BaseDecode(token)
	if claims == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

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
}

func RequireAdmin(c *gin.Context) {
	session := sessions.Default(c)

	token, ok := utils.GetTokenFromSession(session)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not present in the request",
		})
	}

	claims, _ := utils.BaseDecode(token)
	if claims == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Check if it's expired
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Find the user via token sub
	var user models.User
	config.GetDB().First(&user, claims["sub"])

	if !user.IsAdmin {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Attach it to the request
	c.Set("user", user)

	// Continue
	c.Next()
}
