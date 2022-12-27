//
//  Store.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import SQLite
import SQLite3
import Foundation

struct Store {
    var db: Connection? = nil
    var fileDocument: RecipeDBFileDocument? = nil
    
    init() {
        do {
            let url = Store.defaultStoreURL()
            db = try Connection(url.absoluteString)
            fileDocument = RecipeDBFileDocument(url: url)
        } catch {
            print("Failed to open database: \(error)")
        }
    }
    
    init(url: URL) {
        do {
            let dest = Store.defaultStoreURL()
            if (url != dest) {
                print("copy \(url) to \(dest), target exists: \(FileManager.default.fileExists(atPath: dest.absoluteString))")
                if FileManager.default.fileExists(atPath: dest.path) {
                    print("delete \(dest)")
                    try FileManager.default.removeItem(at: dest)
                }
                if url.startAccessingSecurityScopedResource() {
                    try FileManager.default.copyItem(at: url, to: dest)
                    url.stopAccessingSecurityScopedResource()
                } else {
                    print("Unable to access given file!")
                }
            }
        } catch {
            print("Failed copy given database: \(error)")
        }
        
        self.init()
    }
    
    static func defaultStoreURL() -> URL {
        if var url = FileManager.default.urls(for: .applicationSupportDirectory, in: .userDomainMask).first {
#if os(macOS)
            url.append(component: "recipeDB")
            if !FileManager.default.fileExists(atPath: url.path) {
                do {
                    try FileManager.default.createDirectory(at: url, withIntermediateDirectories: false)
                } catch {
                    fatalError("Failed to create application directory: \(url)")
                }
            }
#endif
            url.append(path: "recipeDB.db")
            
            return url
        }
        
        fatalError("Unable to determine application directory")
    }
    
    func getRecipes() -> [Recipe] {
        if (db == nil) {
            return [Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe(),Recipe()];
        }
        
        var recipes: [Recipe] = []
        do {
            let qry = Recipe.sqlTable
                .select(Recipe.sqlColumnId, Recipe.sqlColumnName)
                .filter(Recipe.sqlColumnDeletedAt == nil)
                .order(Recipe.sqlColumnName.asc)
            for recipe in try db!.prepare(qry) {
                recipes.append(Recipe(recipe: recipe,
                                      ingredients: getIngredients(recipeId: recipe[Recipe.sqlColumnId]),
                                      steps: getSteps(recipeId: recipe[Recipe.sqlColumnId])))
            }
        } catch {
            print(error)
        }
        
        return recipes
    }
    
    func getIngredients(recipeId: Int64) -> [Ingredient] {
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
    
    func getSteps(recipeId: Int64) -> [Step] {
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
            do {
                try db.transaction {
                    recipe.recordId = try db.run(Recipe.sqlTable.insert(or: .replace,
                                                                        Recipe.sqlColumnId <- recipe.recordId!,
                                                                        Recipe.sqlColumnName <- recipe.name,
                                                                        Recipe.sqlColumnUpdatedAt <- Date()
                                                                       ))
                    try saveIngredients(recipe: recipe)
                    try saveSteps(recipe: recipe)
                }
            } catch {
                print(error)
            }
        }
    }
    
    private func saveIngredients(recipe: Recipe) throws {
        if let db = db {
            try db.run(Ingredient.sqlTable
                .filter(Ingredient.sqlColumnRecipeId == recipe.recordId!)
                .delete())
            for ingredient in recipe.ingredients {
                try db.run(Ingredient.sqlTable
                    .insert(
                        Ingredient.sqlColumnName <- ingredient.name,
                        Ingredient.sqlColumnRecipeId <- recipe.recordId!,
                        Ingredient.sqlColumnAmount <- ingredient.amount,
                        Ingredient.sqlColumnUnit <- ingredient.unit
                    ))
            }
        }
    }
    
    private func saveSteps(recipe: Recipe) throws {
        if let db = db {
            try db.run(Step.sqlTable
                .filter(Step.sqlColumnRecipeId == recipe.recordId!)
                .delete()
            )
            for step in recipe.steps {
                try db.run(Step.sqlTable
                    .insert(
                        Step.sqlColumnRecipeId <- recipe.recordId!,
                        Step.sqlColumnDescription <- step.description
                    ))
            }
        }
    }
    
    public func close() {
        if let db = db {
            sqlite3_close(db.handle)
        }
    }
}
