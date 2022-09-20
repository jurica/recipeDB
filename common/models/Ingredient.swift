//
//  Ingredient.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import Foundation
import SQLite

class Ingredient : Hashable, Identifiable {
    static func == (lhs: Ingredient, rhs: Ingredient) -> Bool {
        return lhs.id == rhs.id
    }
    
    var id: Int
    var recipeId: Int
    var name: String
    var amount: String
    var unit: String
    
    static let sqlTable = Table("ingredients")
    static let sqlColumnId = Expression<Int>("id")
    static let sqlColumnRecipeId = Expression<Int>("recipe_id")
    static let sqlColumnName = Expression<String>("name")
    static let sqlColumnAmount = Expression<String>("amount")
    static let sqlColumnUnit = Expression<String>("unit")
    
    init() {
        self.id = Int.random(in: Int.min...0)
        self.recipeId = Int.random(in: 0...1000)
        self.name = "Mehl"
        self.amount = "500"
        self.unit = "g"
    }
    
    init(recipeId: Int) {
        self.id = Int.random(in: Int.min...0)
        self.recipeId = recipeId
        self.name = ""
        self.amount = ""
        self.unit = ""
    }
    
    init(ingredient: Row) {
        self.id = ingredient[Ingredient.sqlColumnId]
        self.recipeId = ingredient[Ingredient.sqlColumnRecipeId]
        self.name = ingredient[Ingredient.sqlColumnName]
        self.amount = ingredient[Ingredient.sqlColumnAmount]
        self.unit = ingredient[Ingredient.sqlColumnUnit]
    }
    
    func hash(into hasher: inout Hasher) {
        hasher.combine(id)
    }
    
    func amountWithUnit() -> String {
        return "\(amount) \(unit)"
    }
}
