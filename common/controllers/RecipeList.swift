//
//  RecipeList.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 24.09.22.
//

import Foundation

class RecipeList : ObservableObject {
    private var recipes: [Recipe] = []
    private var searchFor: String = ""
    private var store: Store? = Store()
    
    init() {
        if let store {
            recipes = store.getRecipes()
        }
    }
    
    func test() -> RecipeDBFileDocument? {
        return store?.fileDocument
    }
    
    func refresh() {
        if let store = store {
            recipes = store.getRecipes()
            objectWillChange.send()
        }
    }
    
    func save(recipe: Recipe) {
        if let store = store {
            store.save(recipe: recipe)
            refresh()
        }
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
    
    func initStore(url: URL) {
        store = nil
        store = Store(url: url)
        refresh()
    }
}
