package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
)

var db, dbErr = scribble.New("./data", nil)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.Use(static.Serve("/", static.LocalFile("ui", true)))

	// r.GET("/", index)
	r.GET("/recipe", httpGetRecipes)
	r.GET("/recipe/:id", httpGetRecipe)

	r.POST("/recipe", httpPostRecipe)

	r.PUT("/recipe", httpPutRecipe)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func index(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"/recipe": "list all recipes",
	})
}

func httpGetRecipe(c *gin.Context) {
	recipe := Recipe{}
	err := db.Read("recipe", c.Param("id"), &recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, recipe)
	}
}

func httpGetRecipes(c *gin.Context) {
	records, err := db.ReadAll("recipe")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {

		recipes := []Recipe{}
		for _, f := range records {
			recipeFound := Recipe{}
			if err := json.Unmarshal([]byte(f), &recipeFound); err != nil {
				fmt.Println("Error", err)
			} else {
				recipes = append(recipes, recipeFound)
			}
		}

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

	if recipe.ID == uuid.Nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID missing for recipe",
		})
	}

	recipe, err = CreateOrUpdateRecipe(recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

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

	recipe, err = CreateOrUpdateRecipe(recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusAccepted, recipe)
}
