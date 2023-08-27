package main

import (
	"rent-book-management-system/config"
	"rent-book-management-system/models"
)

func InitMigrations() {
	config.Init("../.env")
	config.GetDB().AutoMigrate(&models.User{}, &models.Book{})
}

func main() {
	InitMigrations()
}
