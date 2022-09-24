//
//  RecipeList.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 24.09.22.
//

import Foundation

class RecipeList : ObservableObject {
    var recipes: [Recipe]
    
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
}
