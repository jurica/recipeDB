//
//  ContentView.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 16.09.22.
//

import SwiftUI

struct ContentView: View {
    @ObservedObject private var recipeList: RecipeList = recipeDBApp.recipeList
    
    var body: some View {
        NavigationSplitView {
            List(recipeList.recipes) { recipe in
                RecipeRow(recipe: recipe)
            }
        } detail: {
            if let initialRecipe = recipeList.recipes.first {
                RecipeDetail(recipe: initialRecipe)
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
