package controllers

import (
	"net/http"

	"bacurin.de/recipeDB/backend/models"
	"github.com/gin-gonic/gin"
)

type recipeControllerInterface interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type recipeControllerStruct struct{}

var (
	// Recipe exposed user controller
	Recipe recipeControllerInterface = &recipeControllerStruct{}
)

func (rc *recipeControllerStruct) Get(c *gin.Context) {
	recipe := models.Recipe{}
	result := models.Model.DB().Preload("Steps").Preload("Ingredients").Find(&recipe, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, recipe)
	}
}

func (rc *recipeControllerStruct) GetAll(c *gin.Context) {
	var recipes []models.Recipe
	result := models.Model.DB().Find(&recipes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, recipes)
	}
}

func (rc *recipeControllerStruct) Update(c *gin.Context) {
	var recipe models.Recipe
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

	tx := models.Model.DB().Begin()

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

	if err = tx.Delete(models.Ingredient{}, "recipe_id = ?", recipe.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = tx.Delete(models.Step{}, "recipe_id = ?", recipe.ID).Error; err != nil {
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

func (rc *recipeControllerStruct) Create(c *gin.Context) {
	var recipe models.Recipe
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

	result := models.Model.DB().Create(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusAccepted, recipe)
}

func (rc *recipeControllerStruct) Delete(c *gin.Context) {
	recipe := models.Recipe{}
	result := models.Model.DB().Find(&recipe, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	result = models.Model.DB().Delete(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, recipe)
}
