//
//  recipe.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 16.09.22.
//

import Foundation
import SQLite

class Recipe : ObservableObject, Hashable, Identifiable {
    // MARK: member variables
    let id: UUID = UUID()
    var recordId: Int64? = nil
    @Published var name: String
    @Published var ingredients: [Ingredient]
    @Published var steps: [Step]
    
    // MARK: sql model
    static let sqlTable = Table("recipes")
    static let sqlColumnId = Expression<Int64>("id")
    static let sqlColumnCreatedAt = Expression<Date?>("created_at")
    static let sqlColumnUpdatedAt = Expression<Date?>("updated_at")
    static let sqlColumnDeletedAt = Expression<Date?>("deleted_at")
    static let sqlColumnName = Expression<String>("name")
    
    // MARK: protocol implementations
    static func == (lhs: Recipe, rhs: Recipe) -> Bool {
        if (lhs.recordId != nil && rhs.recordId != nil) {
            return lhs.recordId == rhs.recordId
        }
        return lhs.id == rhs.id
    }
    
    func hash(into hasher: inout Hasher) {
        hasher.combine(id)
    }
    
    // MARK: initializers
    init() {
        self.name = "Rezept ..."
        self.ingredients = [Ingredient(),Ingredient()]
        self.steps = [Step(),Step()]
    }
    
    init(recipe: Row, ingredients: [Ingredient], steps: [Step]) {
        self.recordId = recipe[Recipe.sqlColumnId]
        self.name = recipe[Recipe.sqlColumnName]
        self.ingredients = ingredients
        self.steps = steps
    }
    
    // MARK: functions
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
            steps.insert(Step(), at: idx+1)
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
            ingredients.insert(Ingredient(), at: idx+1)
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
