//
//  Step.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import Foundation
import SQLite

class Step : Hashable, Identifiable {
    static func == (lhs: Step, rhs: Step) -> Bool {
        return lhs.id == rhs.id
    }
    
    var id: Int?
    var recipeId: Int
    var description: String
    
    static let sqlTable = Table("steps")
    static let sqlColumnId = Expression<Int>("id")
    static let sqlColumnRecipeId = Expression<Int>("recipe_id")
    static let sqlColumnDescription = Expression<String>("description")
    
    init() {
        self.recipeId = Int.random(in: Int.min...0)
        self.description = ""
    }
    
    init(recipeId: Int) {
        self.recipeId = recipeId
        self.description = ""
    }
    
    init(step: Row) {
        self.id = step[Step.sqlColumnId]
        self.recipeId = step[Step.sqlColumnRecipeId]
        self.description = step[Step.sqlColumnDescription]
    }
    
    func hash(into hasher: inout Hasher) {
        hasher.combine(id)
    }
}
