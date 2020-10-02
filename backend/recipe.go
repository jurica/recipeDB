package main

import (
	"gorm.io/gorm"
)

// Recipe instruction how to cook/bake/make something
type Recipe struct {
	gorm.Model
	Name        string
	Ingredients []Ingredient //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Steps       []Step       //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Ingredient single ingredient of a recipe
type Ingredient struct {
	ID       uint
	RecipeID uint
	Name     string
	Amount   string
	Unit     string
}

// Step contains the description text
type Step struct {
	ID          uint
	RecipeID    uint
	Description string
}
