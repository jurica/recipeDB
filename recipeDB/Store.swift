//
//  Store.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import SQLite

struct Store {
    var db: Connection? = nil
    
    init() {
        do {
            db = try Connection("recipeDB.db")
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
}
