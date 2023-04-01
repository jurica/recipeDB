package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"bacurin.de/recipeDB/backend/dtos"
	"bacurin.de/recipeDB/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type recipeControllerInterface interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	ExportToMarkdown()
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
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 || limit > 50 {
		limit = 5
	}

	searchQuery := c.Query("searchQuery")

	// var recipes []models.Recipe
	data := dtos.RecipeList{}

	var qryResult *gorm.DB
	if searchQuery != "" {
		searchQuery = "%" + searchQuery + "%"
		qryResult = models.Model.DB().Offset(offset).Limit(limit).Order(c.DefaultQuery("order", "id asc")).Where("name LIKE ?", searchQuery).Find(&data.Recipes)

		models.Model.DB().Model(&models.Recipe{}).Where("name LIKE ?", searchQuery).Count(&data.RecipeCount)
	} else {
		qryResult = models.Model.DB().Offset(offset).Limit(limit).Order(c.DefaultQuery("order", "id asc")).Find(&data.Recipes)
		models.Model.DB().Model(&models.Recipe{}).Count(&data.RecipeCount)
	}

	if qryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": qryResult.Error.Error(),
		})
	} else {
		data.Offset = int64(offset)
		data.Limit = int64(limit)
		data.CurrentPage = (data.Offset / data.Limit) + 1
		data.PageCount = (data.RecipeCount / data.Limit) + 1
		c.JSON(http.StatusOK, data)
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

func (rc *recipeControllerStruct) ExportToMarkdown() {
	data := dtos.RecipeList{}
	var qryResult *gorm.DB
	qryResult = models.Model.DB().Preload("Steps").Preload("Ingredients").Find(&data.Recipes)
	if qryResult.Error == nil {
		models.Model.DB().Model(&models.Recipe{}).Count(&data.RecipeCount)
		println("found", data.RecipeCount, "recipes")

		for _, recipe := range data.Recipes {
			var markdownRecipe strings.Builder
			markdownRecipe.WriteString("# ")
			markdownRecipe.WriteString(recipe.Name)
			markdownRecipe.WriteString("\n\n")

			markdownRecipe.WriteString("## Zutaten\n")
			markdownRecipe.WriteString("| Menge | Einheit | Zutat |\n")
			markdownRecipe.WriteString("| ---: | ---: | --- |\n")
			for _, ingredient := range recipe.Ingredients {
				markdownRecipe.WriteString("| ")
				markdownRecipe.WriteString(ingredient.Amount)
				markdownRecipe.WriteString(" | ")
				markdownRecipe.WriteString(ingredient.Unit)
				markdownRecipe.WriteString(" | ")
				markdownRecipe.WriteString(ingredient.Name)
				markdownRecipe.WriteString(" |\n")
			}
			markdownRecipe.WriteString("\n")

			markdownRecipe.WriteString("## Zubereitung\n")
			for i, step := range recipe.Steps {
				markdownRecipe.WriteString("### Schritt ")
				markdownRecipe.WriteString(strconv.FormatInt(int64(i+1), 10))
				markdownRecipe.WriteString("\n\n")
				markdownRecipe.WriteString(step.Description)
				markdownRecipe.WriteString("\n\n")
			}

			// println(markdownRecipe.String())

			f, err := os.Create("md/"+recipe.Name+".md")
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			_, err2 := f.WriteString(markdownRecipe.String())
			if err2 != nil {
				log.Fatal(err2)
			}

			// return
		}
	}
}