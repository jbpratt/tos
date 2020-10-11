//
//  SettingsViewModel.swift
//  tos
//
//  Created by John Pratt on 10/10/20.
//

import Foundation

final class SettingsViewModel: ObservableObject, Identifiable {
    // TODO: move to whatever place ios wants strings defined
    @Published var appTitle: String = "TOS"
}
