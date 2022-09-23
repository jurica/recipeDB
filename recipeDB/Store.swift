//
//  Store.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import SQLite
import Foundation

struct Store {
    var db: Connection? = nil
    
    init() {
        do {
//            print("ubiquityIdentityToken: %@", FileManager.default.ubiquityIdentityToken as Any)
            if let path = FileManager.default.url(forUbiquityContainerIdentifier: "iCloud.de.bacurin.recipeDB") {
                let documents = path.appendingPathComponent("Documents")
                let dbFile = documents.appending(path: "recipeDB.db")
//                print(FileManager.default.isUbiquitousItem(at: dbFile))
                db = try Connection(dbFile.absoluteString)
            }
        } catch {
            print(error)
        }
    }
    
    func getRecipes() -> [Recipe] {
        if (db == nil) {
            return [Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe()];
        }
        
        let id = Expression<Int>("id")
        let name = Expression<String>("name")
        
        var recipes: [Recipe] = []
        do {
            let tblRecipes = Table("recipes")
            for recipe in try db!.prepare(tblRecipes) {
                recipes.append(Recipe(id: recipe[id], name: recipe[name], ingredients: getIngredients(recipeId: recipe[id]), steps: getSteps(recipeId: recipe[id])))
            }
        } catch {
            print(error)
        }
        
        return recipes
    }
    
    func getIngredients(recipeId: Int) -> [Ingredient] {
        var ingredients: [Ingredient] = []
        
        if (db != nil) {
            do {
                for ingredient in try db!.prepare(Ingredient.sqlTable.filter(Ingredient.sqlColumnRecipeId == recipeId)) {
                    ingredients.append(Ingredient(ingredient: ingredient))
                }
            } catch {
                print(error)
            }
        }
        
        return ingredients
    }
    
    func getSteps(recipeId: Int) -> [Step] {
        var steps: [Step] = []
        
        if (db != nil) {
            do {
                for step in try db!.prepare(Step.sqlTable.filter(Step.sqlColumnRecipeId == recipeId)) {
                    steps.append(Step(step: step))
                }
            } catch {
                print(error)
            }
        }
        
        return steps
    }
    
    func save(recipe: Recipe) {
        if let db = db {
//            try db.run(users.insert(or: .replace, email <- "alice@mac.com", name <- "Alice B."))
            do {
                try db.transaction {
                    try db.run(Recipe.sqlTable.insert(or: .replace,
                                                      Recipe.sqlColumnId <- recipe.id,
                                                      Recipe.sqlColumnName <- recipe.name))
                }
            } catch {
                print(error)
            }
        }
    }
}
