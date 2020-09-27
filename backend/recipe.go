package main

import (
	"gorm.io/gorm"
)

// Recipe instruction how to cook/bake/make something
type Recipe struct {
	gorm.Model
	Name        string
	Ingredients []Ingredient
	Steps       []Step
}

// Ingredient single ingredient of a recipe
type Ingredient struct {
	gorm.Model
	RecipeID uint
	Name     string
	Amount   string
	Unit     string
}

// Step contains the description text
type Step struct {
	gorm.Model
	RecipeID    uint
	Description string
}

// // NewRecipe get a new recipe with unique ID
// func NewRecipe() Recipe {
// 	recipe := Recipe{
// 		ID: uuid.New(),
// 	}

// 	return recipe
// }

// CreateOrUpdateRecipe persists a recipe
// func CreateOrUpdateRecipe(recipe Recipe) (newRecipe Recipe, err error) {
// 	if recipe.ID == uuid.Nil {
// 		recipe.ID = uuid.New()
// 	}

// 	// TODO couchDB like check for changes
// 	if err := dbJSON.Write("recipe", recipe.ID.String(), recipe); err != nil {
// 		return recipe, err
// 	}

// 	return recipe, nil
// }

// DeleteRecipe delete a recipe
func DeleteRecipe(recipeID string) (err error) {
	return dbJSON.Delete("recipe", recipeID)
}
