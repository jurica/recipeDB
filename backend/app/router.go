package app

import (
	"bacurin.de/recipeDB/backend/controllers"
	"bacurin.de/recipeDB/backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func route() {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	r.GET("/recipelist", middlewares.TokenAuthMiddleware(), controllers.Recipe.GetAll)
	r.GET("/recipe/:id", middlewares.TokenAuthMiddleware(), controllers.Recipe.Get)

	r.POST("/recipe", middlewares.TokenAuthMiddleware(), controllers.Recipe.Update)

	r.PUT("/recipe", middlewares.TokenAuthMiddleware(), controllers.Recipe.Create)

	r.DELETE("/recipe/:id", middlewares.TokenAuthMiddleware(), controllers.Recipe.Delete)

	r.POST("/login", controllers.User.Login)
	r.POST("/refresh-token", middlewares.TokenAuthMiddleware(), controllers.User.RefreshToken)
}
