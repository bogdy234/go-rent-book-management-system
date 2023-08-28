package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Books     []*Book `gorm:"many2many:user_books"`
	IsAdmin   bool    `json:"isAdmin" gorm:"default:false"`
}

type CreateUserInput struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required,min=3,max=50"`
	LastName  string `json:"lastName" validate:"required,min=3,max=50"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
