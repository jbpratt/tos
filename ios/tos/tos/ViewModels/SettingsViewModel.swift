import Foundation

final class SettingsViewModel: ObservableObject, Identifiable {
    // TODO: move to whatever place ios wants strings defined
    @Published var appTitle: String = "TOS"
}
