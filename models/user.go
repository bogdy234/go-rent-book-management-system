package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string
	Books []*Book `gorm:"many2many:user_books"`
}
