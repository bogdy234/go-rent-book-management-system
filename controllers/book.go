package controllers

import (
	"net/http"
	"rent-book-management-system/config"
	"rent-book-management-system/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Body struct {
	Title       string
	Author      string
	Description string
	PricePerDay float32
	Qty         int
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

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

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	// validate book
	err := validate.Struct(book)
	if err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)

		c.JSON(http.StatusBadRequest, gin.H{"errors": errs.Translate(trans)})
		return
	}

	if err := config.GetDB().Create(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	// id, _ := c.Params.Get("id")

	// var b Body
	// c.Bind(&b)

	// var book models.Book
	// config.GetDB().First(&book, id).Update()

	// if book.ID == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Book with provided ID not found"})
	// 	return
	// }

	// c.JSON(200, gin.H{
	// 	"book": book,
	// })
}

func DeleteBook(c *gin.Context) {

}
