package main

import (
	"rent-book-management-system/config"
	"rent-book-management-system/routes"
)

func main() {
	config.Init()
	r := routes.SetupRouter()

	r.Run(":8080")
}
