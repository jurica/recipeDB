//
//  DetailView.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 23.12.22.
//

import SwiftUI
import UniformTypeIdentifiers

struct RecipeToolbar: ToolbarContent {
    let toggleFileImporter: () -> Void
    let toggleFileExporter: () -> Void
    
    var body: some ToolbarContent {
        ToolbarItem(content: {
            Button(action: {
                toggleFileExporter()
            }) {
                Image(systemName: "icloud.and.arrow.up")
                    .imageScale(.large)
            }
            .buttonStyle(MyButtonStyle())
        })
        ToolbarItem(content: {
            Button(action: {
                toggleFileImporter()
            }) {
                Image(systemName: "icloud.and.arrow.down")
                    .imageScale(.large)
            }
            .buttonStyle(MyButtonStyle())
        })
        ToolbarItem(content: {
            Button(action: {
                print("nop")
            }) {
                Image(systemName: "square.and.pencil")
                    .imageScale(.large)
            }
            .buttonStyle(MyButtonStyle())
        })
    }
}

struct DetailView: View {
    @ObservedObject var recipe: Recipe
    @State private var settingsSheetIsPresented = false
    @State var fileImporterIsPresented = false
    @State var fileExporterIsPresented = false
    
    var body: some View {
        RecipeDetail(recipe: recipe)
            .toolbar{
                RecipeToolbar(toggleFileImporter: {
                    fileImporterIsPresented.toggle()
                }, toggleFileExporter: {
                    fileExporterIsPresented.toggle()
                })
            }
            .fileImporter(isPresented: $fileImporterIsPresented, allowedContentTypes: [.item], allowsMultipleSelection: false) { result in
                do {
                    let url = try result.get().first!
                    recipeDBApp.recipeList.initStore(url: url)
                } catch {
                    print("Error: \(error)")
                }
            }
            .fileExporter(isPresented: $fileExporterIsPresented, document: recipeDBApp.recipeList.test(), contentType: UTType.database) { result in
                switch result {
                case .success(let url):
                    print("saved to \(url)")
                case.failure(let error):
                    print(error.localizedDescription)
                }
            }
    }
}
