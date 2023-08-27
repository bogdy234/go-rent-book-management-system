package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	PricePerDay float32 `json:"pricePerDay"`
	Users       []*User `gorm:"many2many:user_books"`
	Qty         int     `json:"qty"`
}

type CreateBookInput struct {
	Title       string  `json:"title" validate:"required,min=4,max=50"`
	Author      string  `json:"author" validate:"required,min=2"`
	Description string  `json:"description" validate:"required,min=10"`
	PricePerDay float32 `json:"pricePerDay" validate:"min=0"`
	Users       []*User `gorm:"many2many:user_books"`
	Qty         int     `json:"qty" validate:"min=0"`
}

type UpdateBookInput struct {
	Title       string  `json:"title" validate:"omitempty,min=4,max=50"`
	Author      string  `json:"author" validate:"omitempty,min=2"`
	Description string  `json:"description" validate:"omitempty,min=10"`
	PricePerDay float32 `json:"pricePerDay" validate:"omitempty,min=0"`
	Users       []*User `gorm:"many2many:user_books"`
	Qty         int     `json:"qty" validate:"omitempty,min=0"`
}
