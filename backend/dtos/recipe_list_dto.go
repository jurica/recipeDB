package dtos

import "bacurin.de/recipeDB/backend/models"

type RecipeList struct {
	RecipeCount int64
	Limit       int64
	Offset      int64
	PageCount   int64
	CurrentPage int64
	Recipes     []models.Recipe
}
