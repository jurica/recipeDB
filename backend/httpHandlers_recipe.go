package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		return
	}

	if recipe.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID missing for recipe",
		})
	}

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = tx.Delete(Ingredient{}, "recipe_id = ?", recipe.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = tx.Delete(Step{}, "recipe_id = ?", recipe.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = tx.Save(&recipe).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx.Commit()

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

	if recipe.ID != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID given for new recipe",
		})
	}

	result := db.Create(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusAccepted, recipe)
}

func httpDeleteRecipe(c *gin.Context) {
	recipe := Recipe{}
	result := db.Find(&recipe, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	result = db.Delete(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusNoContent, nil)
}