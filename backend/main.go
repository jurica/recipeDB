package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// TODO configure path to database file
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

	r.GET("/recipe", httpGetRecipes)
	r.GET("/recipe/:id", httpGetRecipe)

	r.POST("/recipe", httpPostRecipe)

	r.PUT("/recipe", httpPutRecipe)

	r.DELETE("/recipe/:id", httpDeleteRecipe)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
