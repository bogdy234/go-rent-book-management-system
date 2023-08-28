package controllers

import (
	"net/http"
	"rent-book-management-system/config"
	"rent-book-management-system/models"
	"rent-book-management-system/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
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
	session := sessions.Default(c)

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

	// encode data in JWT and add it to session
	tokenString := utils.EncodeDataJWT(jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	session.Set("token", tokenString)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"email":     user.Email,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
		},
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()
	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}

func UpdateUser(c *gin.Context) {
	var u models.UpdateUser
	c.Bind(&u)

	// validate user
	err := config.Validate.Struct(u)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs.Translate(config.Trans)})
		return
	}

	// get user from Context
	userReq, _ := c.Get("user")
	user, _ := userReq.(models.User)

	// update user
	if err := config.GetDB().Model(&user).Updates(u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return new user data
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	// get user from Context
	userReq, _ := c.Get("user")
	user, _ := userReq.(models.User)

	config.GetDB().Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
