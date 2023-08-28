package routes

import (
	"net/http"
	"rent-book-management-system/config"
	"rent-book-management-system/constants"
	"rent-book-management-system/controllers"
	"rent-book-management-system/middlewares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("user_session", config.SessionStore))
	// r.Use()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Book CRUD
	r.POST(constants.BookRoute, middlewares.RequireAdmin, controllers.CreateBook)
	r.GET(constants.BookIdRoute, controllers.GetBook)
	r.PUT(constants.BookIdRoute, controllers.UpdateBook)
	r.DELETE(constants.BookIdRoute, controllers.DeleteBook)

	// User
	r.POST(constants.UserRoute, controllers.CreateUser)
	r.POST(constants.LoginRoute, controllers.Login)
	r.POST(constants.LogoutRoute, controllers.Logout)
	r.PUT(constants.UserRoute, middlewares.RequireAuth, controllers.UpdateUser)
	r.DELETE(constants.UserRoute, middlewares.RequireAuth, controllers.DeleteUser)

	// Test session
	r.GET("/testsession", controllers.TestSession)
	r.GET("/test-jwt", controllers.TestJwt)

	return r
}
