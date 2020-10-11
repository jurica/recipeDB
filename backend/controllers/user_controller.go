package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"bacurin.de/recipeDB/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userControllerInterface interface {
	Login(*gin.Context)
	RefreshToken(*gin.Context)
	HashAndSaltPassword()
}

type userControllerStruct struct{}

var (
	// User exposed user controller
	User userControllerInterface = &userControllerStruct{}
)

func (us *userControllerStruct) Login(c *gin.Context) {
	var user models.User
	var err error

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var plainPassword = user.Password
	models.Model.DB().Where("email = ?", user.Email).First(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainPassword))

	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "fail!",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_uuid"] = uuid.NewV4().String()
	claims["user_id"] = 1
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("superDuperSaveRecipeDB"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	user.Token = tokenStr

	c.JSON(http.StatusOK, user)
}

func (us *userControllerStruct) RefreshToken(c *gin.Context) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_uuid"] = uuid.NewV4().String()
	claims["user_id"] = 1
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("superDuperSaveRecipeDB"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	user := models.User{}
	user.ID = 1
	user.Token = tokenStr

	c.JSON(http.StatusOK, user)
}

func (us *userControllerStruct) HashAndSaltPassword() {
	fmt.Println("Enter password")
	var password string
	_, err := fmt.Scan(&password)
	if err != nil {
		log.Println(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Plain password: ", password)
	fmt.Println("Salted Hash: ", string(hash))

	test := []byte("$2a$10$85G75zTBk7.6hP3DfKuQvuCeLFc3ZLR7Yp52EbuOXU0D/6IdOq0zO")
	err = bcrypt.CompareHashAndPassword(test, hash)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("password match!")
}
