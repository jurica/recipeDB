//
//  ButtonStyle.swift
//  recipeDB
//
//  Created by Jurica Bacurin on 17.09.22.
//

import SwiftUI

struct MyButtonStyle: ButtonStyle {
    @State var isHovered: Bool = false
    
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .padding(3)
            .onHover { hovered in
                self.isHovered = hovered
            }
            .background(
                ZStack{
                    if isHovered {
                        RoundedRectangle(cornerRadius: 5)
                            .fill(Color.primary.opacity(0.1))
                    } else {
                        RoundedRectangle(cornerRadius: 5)
                            .fill(.clear)
                    }
                }
            )
    }
}
