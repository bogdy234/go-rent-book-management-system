package controllers

import (
	"fmt"
	"net/http"
	"os"
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

// TEST

func TestSession(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("hello") != "world" {
		session.Set("hello", "world")
		session.Save()
	}

	// session.Options({

	// })
	// session.Set("token", "MTY5MzE2MzEwNHxEdi1CQkFFQ180SUFBUkFCRUFBQUpQLUNBQUVHYzNSeWFXNW5EQWNBQldobGJHeHZCbk4wY21sdVp3d0hBQVYzYjNKc1pBPT18utaIWdDjycs7rOSQvpAdp8YhS61VgxhdNthe05Z6nRw%3D")

	c.JSON(http.StatusOK, gin.H{"session": session.Get("token")})
}

func TestJwt(c *gin.Context) {
	byteArrSecret := []byte(os.Getenv("JWT_SECRET"))

	// ENCODE
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(byteArrSecret)

	fmt.Println(tokenString, err)

	// DECODE
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return byteArrSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
