//
//  recipe.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 16.09.22.
//

import Foundation
import SQLite

class Recipe : ObservableObject, Hashable {
    static func == (lhs: Recipe, rhs: Recipe) -> Bool {
        return lhs.id == rhs.id
    }
    
    var id: Int?
    @Published var name: String
    @Published var ingredients: [Ingredient]
    @Published var steps: [Step]
    
    static let sqlTable = Table("recipes")
    static let sqlColumnId = Expression<Int?>("id")
    static let sqlColumnName = Expression<String?>("name")
    
    init() {
        self.name = "Rezept ..."
        self.ingredients = [Ingredient(),Ingredient()]
        self.steps = [Step(),Step()]
    }
    
    init(id: Int, name: String, ingredients: [Ingredient], steps: [Step]) {
        self.id = id
        self.name = name
        self.ingredients = ingredients
        self.steps = steps
    }
    
    func hash(into hasher: inout Hasher) {
        hasher.combine(id)
    }
    
    func stepNumber(step: Step) -> Int {
        let idx = steps.firstIndex(of: step) ?? 0
        return idx + 1
    }
    
    func remove(step: Step) {
        if let idx = steps.firstIndex(of: step) {
            steps.remove(at: idx)
        }
    }
    
    func add(after: Step) {
        if let idx = steps.firstIndex(of: after) {
            if let id = id {
                steps.insert(Step(recipeId: id), at: idx+1)
            }
        }
    }
    
    func moveUp(step: Step) {
        if let idx = steps.firstIndex(of: step) {
            if (idx > 0) {
                steps.swapAt(idx-1, idx)
            }
        }
    }
    
    func moveDown(step: Step) {
        if let idx = steps.firstIndex(of: step) {
            if (idx < steps.count-1) {
                steps.swapAt(idx+1, idx)
            }
        }
    }
    
    func remove(ingredient: Ingredient) {
        if let idx = ingredients.firstIndex(of: ingredient) {
            ingredients.remove(at: idx)
        }
    }
    
    func add(after: Ingredient) {
        if let idx = ingredients.firstIndex(of: after) {
            if let id = id {
                ingredients.insert(Ingredient(recipeId: id), at: idx+1)
            }
        }
    }
    
    func moveUp(ingredient: Ingredient) {
        if let idx = ingredients.firstIndex(of: ingredient) {
            if (idx > 0) {
                ingredients.swapAt(idx-1, idx)
            }
        }
    }
    
    func moveDown(ingredient: Ingredient) {
        if let idx = ingredients.firstIndex(of: ingredient) {
            if (idx < ingredients.count-1) {
                ingredients.swapAt(idx+1, idx)
            }
        }
    }
}
