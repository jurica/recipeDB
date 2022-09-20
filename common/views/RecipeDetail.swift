//
//  RecipeRow.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import SwiftUI

struct RecipeDetail: View {
    @ObservedObject var recipe: Recipe
    @State var editing: Bool = false
    var body: some View {
        VStack {
            ScrollView {
                if editing {
                    TextField("" , text: Binding(
                        get: {recipe.name},
                        set: {recipe.name = $0
                            recipe.objectWillChange.send()
                        }
                    ))
                    .textFieldStyle(.roundedBorder)
                } else {
                    HStack {
                        Text(recipe.name)
                            .fontWeight(.heavy)
                            .font(.title)
                        Spacer()
                        Button(action: {
                            print("nop")
                        }) {
                            Image(systemName: "trash")
                                .imageScale(.large)
                        }
                        .buttonStyle(MyButtonStyle())
                        Button(action: {
                            editing = true
                        }) {
                            Image(systemName: "pencil")
                                .imageScale(.large)
                        }
                        .buttonStyle(MyButtonStyle())
                    }
                }
                if editing {
                    Text("Zutaten")
                        .font(.title)
                    ForEach(recipe.ingredients) { ingredient in
                        HStack {
                            TextField("Zutat", text: Binding(get: {
                                ingredient.name
                            }, set: {
                                ingredient.name = $0
                            }))
                            .textFieldStyle(.roundedBorder)
                            TextField("Menge", text: Binding(get: {
                                ingredient.amount
                            }, set: {
                                ingredient.amount = $0
                            }))
                            .textFieldStyle(.roundedBorder)
                            .frame(width: 75)
                            TextField("Einheit", text: Binding(get: {
                                ingredient.unit
                            }, set: {
                                ingredient.unit = $0
                            }))
                            .textFieldStyle(.roundedBorder)
                            .frame(width: 75)
                            GroupBox {
                                HStack {
                                    Button(action: {
                                        recipe.remove(ingredient: ingredient)
                                    }) {
                                        Image(systemName: "minus")
                                            .imageScale(.large)
                                    }
                                    .buttonStyle(MyButtonStyle())
                                    Button(action: {
                                        recipe.moveUp(ingredient: ingredient)
                                    }) {
                                        Image(systemName: "chevron.up")
                                            .imageScale(.large)
                                    }
                                    .buttonStyle(MyButtonStyle())
                                    Button(action: {
                                        recipe.moveDown(ingredient: ingredient)
                                    }) {
                                        Image(systemName: "chevron.down")
                                            .imageScale(.large)
                                    }
                                    .buttonStyle(MyButtonStyle())
                                    Button(action: {
                                        recipe.add(after: ingredient)
                                    }) {
                                        Image(systemName: "plus")
                                            .imageScale(.large)
                                    }
                                    .buttonStyle(MyButtonStyle())
                                }
                            }
                        }
                    }
                } else {
                    HStack {
                        Image(systemName: "camera")
                            .resizable()
                            .scaledToFit()
                            .frame(width: 100, height: 100)
                            .frame(maxWidth: .infinity)
                        VStack {
                            GroupBox{
                                Text("Zutaten")
                                    .font(.title)
                                    .padding(.bottom)
                                ForEach(recipe.ingredients) { ingredient in
                                    HStack {
                                        Text(ingredient.name)
                                            .fontWeight(.bold)
                                        Spacer()
                                        Text("\(ingredient.amount) \(ingredient.unit)")
                                            .fontWeight(.bold)
                                    }
                                }
                            }
                            Spacer()
                        }
                        .frame(maxWidth: .infinity)
                    }
                }
                Text("Zubereitung")
                    .font(.title)
                    .padding(.bottom)
                if editing {
                    ForEach(recipe.steps) { step in
                        VStack {
                            HStack {
                                Text("Schritt \(recipe.stepNumber(step: step))")
                                    .font(.headline)
                                Spacer()
                                GroupBox {
                                    HStack {
                                        Button(action: {
                                            recipe.remove(step: step)
                                        }) {
                                            Image(systemName: "minus")
                                                .imageScale(.large)
                                        }
                                        .buttonStyle(MyButtonStyle())
                                        Button(action: {
                                            recipe.moveUp(step: step)
                                        }) {
                                            Image(systemName: "chevron.up")
                                                .imageScale(.large)
                                        }
                                        .buttonStyle(MyButtonStyle())
                                        Button(action: {
                                            recipe.moveDown(step: step)
                                        }) {
                                            Image(systemName: "chevron.down")
                                                .imageScale(.large)
                                        }
                                        .buttonStyle(MyButtonStyle())
                                        Button(action: {
                                            recipe.add(after: step)
                                        }) {
                                            Image(systemName: "plus")
                                                .imageScale(.large)
                                        }
                                        .buttonStyle(MyButtonStyle())
                                    }
                                }
                            }
                            GroupBox {
                                TextEditor(text: Binding(
                                    get: { step.description },
                                    set: { step.description = $0 }
                                ))
                                .multilineTextAlignment(.leading)
                                .scrollContentBackground(.hidden)
                            }
                        }
                    }
                } else {
                    ForEach(recipe.steps) { step in
                        Text("Schritt \(recipe.stepNumber(step: step))")
                            .font(.headline)
                            .frame(maxWidth: .infinity, alignment: .leading)
                        Text(step.description)
                            .textSelection(.enabled)
                            .multilineTextAlignment(.leading)
                            .frame(maxWidth: .infinity, alignment: .leading)
                            .padding(.bottom)
                    }
                }
            }
            .padding()
            .background(in: Rectangle())
            if editing {
                HStack {
                    Spacer()
                    Button(action: {
                        editing = false
                    }) {
                        Image(systemName: "xmark")
                            .imageScale(.large)
                    }
                    .buttonStyle(MyButtonStyle())
                    Button(action: {
                        editing = false
                    }) {
                        Image(systemName: "checkmark")
                            .imageScale(.large)
                    }
                    .buttonStyle(MyButtonStyle())
                }
                .padding()
            }
        }
    }
}

struct RecipeDetail_Previews: PreviewProvider {
    static var previews: some View {
        RecipeDetail(recipe: Recipe())
    }
}
