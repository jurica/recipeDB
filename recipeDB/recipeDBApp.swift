//
//  recipeDBApp.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 16.09.22.
//

import SwiftUI

@main
struct recipeDBApp: App {
    static let recipeList: RecipeList = RecipeList()
    
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}
