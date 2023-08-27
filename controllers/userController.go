package controllers

import (
	"net/http"
	"rent-book-management-system/config"
	"rent-book-management-system/models"
	"rent-book-management-system/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(c *gin.Context) {
	var u models.CreateUserInput
	c.Bind(&u)

	// validate user
	err := config.Validate.Struct(u)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs.Translate(config.Trans)})
		return
	}

	// hash the password
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user := models.User{
		Email:     u.Email,
		Password:  hash,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	// create new entry in database
	if err := config.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(c *gin.Context) {
	var l models.LoginInput
	c.Bind(&l)

	// validate user
	err := config.Validate.Struct(l)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs.Translate(config.Trans)})
		return
	}

	// get user from DB
	var user models.User
	config.GetDB().First(&user, "email = ?", l.Email)

	// check if login data is valid
	validLogin := utils.CheckPasswordHash(l.Password, user.Password)
	if !validLogin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password invalid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"email":     user.Email,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
		},
	})
}
