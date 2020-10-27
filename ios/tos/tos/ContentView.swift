import SwiftUI

struct ContentView: View {
    @ObservedObject var menuViewModel = MenuViewModel()
    @ObservedObject var orderViewModel = OrderViewModel()

    @State private var isSettingsActive: Bool = false

    var body: some View {
        NavigationView {
            GeometryReader { geo in
                HStack {
                    OrderView(viewModel: orderViewModel)
                    Divider()
                    MenuView(
                        menuViewModel: menuViewModel,
                        orderViewModel: orderViewModel)
                        .frame(minWidth: geo.size.width - (geo.size.width / 3))
                }
            }
            .navigationBarTitle("Menu", displayMode: .inline)
            .navigationBarItems(
                trailing: NavigationLink(
                    destination: SettingsView(viewModel: menuViewModel).navigationBarTitle("Settings"),
                    isActive: $isSettingsActive) {
                        Button(action: { isSettingsActive = !isSettingsActive }) {
                            Image(systemName: "gear")
                        }
                }
            )
        }
        .navigationViewStyle(StackNavigationViewStyle())
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
