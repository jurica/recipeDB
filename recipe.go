package main

import (
	uuid "github.com/google/uuid"
)

// Recipe instruction how to cook/bake/make something
type Recipe struct {
	Name        string
	Ingredients []Ingredient
	Steps       []string
	ID          uuid.UUID
}

// NewRecipe get a new recipe with unique ID
func NewRecipe() Recipe {
	recipe := Recipe{
		ID: uuid.New(),
	}

	return recipe
}

// CreateOrUpdateRecipe persists a recipe
func CreateOrUpdateRecipe(recipe Recipe) (newRecipe Recipe, err error) {
	if recipe.ID == uuid.Nil {
		recipe.ID = uuid.New()
	}

	// TODO couchDB like check for changes
	if err := db.Write("recipe", recipe.ID.String(), recipe); err != nil {
		return recipe, err
	}

	return recipe, nil
}
