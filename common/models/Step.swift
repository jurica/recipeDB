//
//  Step.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import Foundation
import SQLite

class Step : Hashable, Identifiable {
    // MARK: member variables
    let id: UUID = UUID()
    var description: String
    
    // MARK: sql model
    static let sqlTable = Table("steps")
    static let sqlColumnId = Expression<Int>("id")
    static let sqlColumnRecipeId = Expression<Int64>("recipe_id")
    static let sqlColumnDescription = Expression<String>("description")
    
    // MARK: protocol implementations
    static func == (lhs: Step, rhs: Step) -> Bool {
        return lhs.id == rhs.id
    }
    
    func hash(into hasher: inout Hasher) {
        hasher.combine(id)
    }
    
    // MARK: initializers
    init() {
        self.description = ""
    }
    
    init(step: Row) {
        self.description = step[Step.sqlColumnDescription]
    }
}
