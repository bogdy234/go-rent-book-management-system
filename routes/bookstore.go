package routes

import (
	"net/http"
	"rent-book-management-system/constants"
	"rent-book-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Book CRUD
	r.POST(constants.BookRoute, controllers.CreateBook)
	r.GET(constants.BookIdRoute, controllers.GetBook)
	r.PUT(constants.BookIdRoute, controllers.UpdateBook)
	r.DELETE(constants.BookIdRoute, controllers.DeleteBook)

	// User CRUD
	r.POST(constants.UserRoute, controllers.CreateUser)
	r.POST(constants.LoginRoute, controllers.Login)

	return r
}
