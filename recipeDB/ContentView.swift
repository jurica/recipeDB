//
//  ContentView.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 16.09.22.
//

import SwiftUI

struct ContentView: View {
    @ObservedObject private var recipeList: RecipeList = recipeDBApp.recipeList
    @State private var searchText: String = ""
    
    
    var body: some View {
        NavigationSplitView {
            List(recipeList.getAll()) { recipe in
                RecipeRow(recipe: recipe)
            }
        } detail: {
            if let initialRecipe = recipeList.getAll().first {
                DetailView(recipe: initialRecipe)
            } else {
                DetailView(recipe: Recipe())
            }
        }
        .searchable(text: $searchText)
        .onChange(of: searchText) { _ in
            recipeList.search(name: searchText)
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
