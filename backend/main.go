package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbJSON, dbErr = scribble.New("./data", nil)
var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open(sqlite.Open("recipeDB.db"), &gorm.Config{})

	if err != nil {
		log.Panic("failed to open database")
	}

	err = db.AutoMigrate(&Recipe{}, &Ingredient{}, &Step{})
	if err != nil {
		log.Panic(err.Error())
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	r.Use(static.Serve("/", static.LocalFile("ui", true)))

	// r.GET("/", index)
	r.GET("/recipe", httpGetRecipes)
	r.GET("/recipe/:id", httpGetRecipe)

	r.POST("/recipe", httpPostRecipe)

	r.PUT("/recipe", httpPutRecipe)

	r.DELETE("/recipe/:id", httpDeleteRecipe)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func index(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"/recipe": "list all recipes",
	})
}

func httpGetRecipe(c *gin.Context) {
	recipe := Recipe{}
	result := db.Preload("Steps").Preload("Ingredients").Find(&recipe, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, recipe)
	}
}

func httpGetRecipes(c *gin.Context) {
	var recipes []Recipe
	result := db.Find(&recipes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, recipes)
	}
}

func httpPostRecipe(c *gin.Context) {
	var recipe Recipe
	var err error

	err = c.BindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if recipe.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID missing for recipe",
		})
	}

	// recipe, err = CreateOrUpdateRecipe(recipe)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }

	c.JSON(http.StatusAccepted, recipe)
}

func httpPutRecipe(c *gin.Context) {
	var recipe Recipe
	var err error

	err = c.BindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// recipe, err = CreateOrUpdateRecipe(recipe)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }

	c.JSON(http.StatusAccepted, recipe)
}

func httpDeleteRecipe(c *gin.Context) {
	var err error

	err = DeleteRecipe(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusNoContent, nil)
}
