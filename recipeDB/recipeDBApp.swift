//
//  recipeDBApp.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 16.09.22.
//

import SwiftUI

@main
struct recipeDBApp: App {
    static let store: Store = Store()
    
    var body: some Scene {
        WindowGroup {
            ContentView(recipes: recipeDBApp.getStore().getRecipes())
        }
    }
    
    static func getStore() -> Store {
        return store
    }
}
