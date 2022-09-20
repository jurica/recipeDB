//
//  RecipeRow.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 16.09.22.
//

import SwiftUI

struct RecipeRow: View {
    var recipe: Recipe
    var body: some View {
        NavigationLink(recipe.name) {
            RecipeDetail(recipe: recipe)
        }
    }
}

struct RecipeRow_Previews: PreviewProvider {
    static var previews: some View {
        RecipeRow(recipe: Recipe())
    }
}
