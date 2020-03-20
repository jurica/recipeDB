package main

// Recipe instruction how to cook/bake/make something
type Recipe struct {
	Name        string
	Ingredients []Ingredient
	Steps       []string
}
