package controllers

import (
	"net/http"
	"rent-book-management-system/config"
	"rent-book-management-system/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateBook(c *gin.Context) {
	var b models.CreateBookInput
	c.Bind(&b)

	// validate book
	err := config.Validate.Struct(b)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs.Translate(config.Trans)})
		return
	}

	book := models.Book{
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
		PricePerDay: b.PricePerDay,
		Qty:         b.Qty,
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
	// get id from query params
	id, _ := c.Params.Get("id")

	var book models.Book
	config.GetDB().First(&book, id)

	// error if record is not found in DB
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// return the record
	c.JSON(200, gin.H{
		"book": book,
	})
}

func UpdateBook(c *gin.Context) {
	var b models.UpdateBookInput
	c.Bind(&b)

	// validate book
	err := config.Validate.Struct(b)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs.Translate(config.Trans)})
		return
	}

	// error if record is not found in DB
	var book models.Book
	if err := config.GetDB().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// update entry in DB
	if err := config.GetDB().Model(&book).Updates(b).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return created book
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func DeleteBook(c *gin.Context) {
	id, _ := c.Params.Get("id")

	if err := config.GetDB().Delete(&models.Book{}, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record successfully deleted"})
}
