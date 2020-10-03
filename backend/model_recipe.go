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