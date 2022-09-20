//
//  ContentView.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 16.09.22.
//

import SwiftUI

struct ContentView: View {
//    @State private var searchText = ""
    private let recipes: [Recipe]
    @State private var selectedRecipe: Recipe?
    
    init(recipes: [Recipe]) {
        self.recipes = recipes
    }
        
    var body: some View {
        NavigationSplitView {
            List(recipes, id: \.id, selection: $selectedRecipe) { recipe in
                RecipeRow(recipe: recipe)
            }
        /*} content: {
            List(recipes, id: \.id, selection: $selectedRecipe) { recipe in
                RecipeRow(recipe: recipe)
            }*/
        } detail: {
            VStack {
                Image(systemName: "globe")
                    .imageScale(.large)
                    .foregroundColor(.accentColor)
                Text(selectedRecipe?.name ?? "")
            }
            .padding()
        }
//    .searchable(text: $searchText)
    .toolbar {
        Button(action: {
            print("nop")
        }) {
            Image(systemName: "square.and.pencil")
                .imageScale(.large)
        }
        .buttonStyle(MyButtonStyle())
    }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView(recipes: [Recipe(),Recipe()])
    }
}
