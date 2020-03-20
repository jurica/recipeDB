package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
)

func main() {
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/", index)
	r.GET("/recipe", listRecipes)

	db, err := scribble.New("db", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	recipe := Recipe{
		Name: "Flammkuchenteig",
	}
	if err := db.Write("recipe", "recipe1", recipe); err != nil {
		fmt.Println("Error", err)
	}

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

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func index(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"/recipe": "list all recipes",
	})
}

func listRecipes(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"recipe-1": "Flammkuchenteig",
		"recipe-2": "Pizzateig",
		"recipe-3": "Gulasch",
	})
}
