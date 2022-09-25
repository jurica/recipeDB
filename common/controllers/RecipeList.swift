//
//  RecipeList.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 24.09.22.
//

import Foundation

class RecipeList : ObservableObject {
    private var recipes: [Recipe]
    private var searchFor: String = ""
    private let store: Store = Store()
    
    init() {
        recipes = store.getRecipes()
    }
    
    func refresh() {
        recipes = store.getRecipes()
        objectWillChange.send()
    }
    
    func save(recipe: Recipe) {
        store.save(recipe: recipe)
        refresh()
    }
    
    func search(name: String) {
        searchFor = name
        objectWillChange.send()
    }
    
    func getAll() -> [Recipe] {
        if (searchFor.isEmpty) {
            return recipes
        } else {
            return recipes.filter{$0.name.lowercased().contains(searchFor.lowercased())}
        }
    }
}
