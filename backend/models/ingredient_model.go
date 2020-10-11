package models

// Ingredient single ingredient of a recipe
type Ingredient struct {
	ID       uint
	RecipeID uint
	Name     string
	Amount   string
	Unit     string
}
