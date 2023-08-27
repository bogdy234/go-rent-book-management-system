package controllers

import (
	"net/http"
	"rent-book-management-system/config"
	"rent-book-management-system/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Body struct {
	Title       string
	Author      string
	Description string
	PricePerDay float32
	Qty         int
}

func CreateBook(c *gin.Context) {
	var b Body
	c.Bind(&b)

	book := models.Book{
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
		PricePerDay: b.PricePerDay,
		Qty:         b.Qty,
	}

	// validate book
	err := config.Validate.Struct(book)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs.Translate(config.Trans)})
		return
	}

	// create new entry in database
	if err := config.GetDB().Create(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return created book
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func GetBook(c *gin.Context) {
	id, _ := c.Params.Get("id")

	var book models.Book
	config.GetDB().First(&book, id)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book with provided ID not found"})
		return
	}

	c.JSON(200, gin.H{
		"book": book,
	})
}

func UpdateBook(c *gin.Context) {

}

func DeleteBook(c *gin.Context) {

}
