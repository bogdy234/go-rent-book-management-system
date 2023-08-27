package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string  `json:"title" validate:"required,min=4,max=50"`
	Author      string  `json:"author" validate:"required,min=2"`
	Description string  `json:"description" validate:"required,min=10"`
	PricePerDay float32 `json:"pricePerDay" validate:"min=0"`
	Users       []*User `gorm:"many2many:user_books"`
	Qty         int     `json:"qty" validate:"min=0"`
}
