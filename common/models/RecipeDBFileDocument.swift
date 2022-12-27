//
//  DatabaseFile.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 27.12.22.
//

import SwiftUI
import UniformTypeIdentifiers

//extension UTType {
//    static let recipeDB = UTType(exportedAs: "de.bacurin.recipeDB")
//}

struct RecipeDBFileDocument: FileDocument {
//    static var readableContentTypes = [UTType.recipeDB]
    static var readableContentTypes = [UTType.database]
    var url: URL
    
    init(url: URL) {
        self.url = url
    }
    
    init(configuration: ReadConfiguration) throws {
        self.url = URL(fileURLWithPath: "")
    }
    
    func fileWrapper(configuration: WriteConfiguration) throws -> FileWrapper {
        let file = try! FileWrapper(url: url, options: .immediate)
        return file
    }
    

}
