package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
)

var db, dbErr = scribble.New("data", nil)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("/Users/jurica.bacurin/Dev/recipeDB/ui", true)))

	// r.GET("/", index)
	r.GET("/recipe", httpGetRecipes)

	r.POST("/recipe", httpPostRecipe)

	r.PUT("/recipe", httpPutRecipe)

	records, err := db.ReadAll("recipe")
	if err != nil {
		fmt.Println("Error", err)
	}

	recipes := []Recipe{}
	for _, f := range records {
		recipeFound := Recipe{}
		if err := json.Unmarshal([]byte(f), &recipeFound); err != nil {
			fmt.Println("Error", err)
		}
		recipes = append(recipes, recipeFound)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func index(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"/recipe": "list all recipes",
	})
}

func httpGetRecipes(c *gin.Context) {
	records, err := db.ReadAll("recipe")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	recipes := []Recipe{}
	for _, f := range records {
		recipeFound := Recipe{}
		if err := json.Unmarshal([]byte(f), &recipeFound); err != nil {
			fmt.Println("Error", err)
		}
		recipes = append(recipes, recipeFound)
	}

	// jsonData := []byte(`[{"Name": "test1", "ID":"12-34"},{"Name": "test1", "ID":"12-34"}`)
	c.JSON(http.StatusAccepted, gin.H{"recipes": recipes})
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
