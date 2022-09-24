//
//  Ingredient.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import Foundation
import SQLite

class Ingredient : Hashable, Identifiable {
    // MARK: member variables
    let id: UUID = UUID()
    var name: String
    var amount: String
    var unit: String
    
    // MARK: sql model
    static let sqlTable = Table("ingredients")
    static let sqlColumnId = Expression<Int64>("id")
    static let sqlColumnRecipeId = Expression<Int64>("recipe_id")
    static let sqlColumnName = Expression<String>("name")
    static let sqlColumnAmount = Expression<String>("amount")
    static let sqlColumnUnit = Expression<String>("unit")
    
    // MARK: protocol implementations
    static func == (lhs: Ingredient, rhs: Ingredient) -> Bool {
        return lhs.id == rhs.id
    }
    
    func hash(into hasher: inout Hasher) {
        hasher.combine(id)
    }
    
    // MARK: initializers
    init() {
        self.name = ""
        self.amount = ""
        self.unit = ""
    }
    
    init(ingredient: Row) {
        self.name = ingredient[Ingredient.sqlColumnName]
        self.amount = ingredient[Ingredient.sqlColumnAmount]
        self.unit = ingredient[Ingredient.sqlColumnUnit]
    }
}
