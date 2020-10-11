package models

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
