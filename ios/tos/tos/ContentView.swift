import SwiftUI

struct ContentView: View {
    @ObservedObject var menuViewModel: MenuViewModel
    @ObservedObject var orderViewModel: OrderViewModel
    @ObservedObject var healthViewModel: HealthViewModel

    @State private var isSettingsActive: Bool = false

    var isMenuServing: Bool {
        // Is there really no one that exports the service name?
        healthViewModel.serviceStatus(healthViewModel.menuServiceName) == .serving
    }

    var body: some View {
        NavigationView {
            GeometryReader { geo in
                HStack {
                    OrderView(vm: orderViewModel)
                    Divider()
                    MenuView(
                        menuViewModel: menuViewModel,
                        orderViewModel: orderViewModel
                    )
                    .frame(minWidth: geo.size.width - (geo.size.width / 3))
                }
            }
            .navigationBarTitle("Menu", displayMode: .inline)
            .navigationBarItems(
                leading: heart,
                trailing: NavigationLink(
                    destination: SettingsView(viewModel: menuViewModel)
                        .navigationBarTitle("Settings"),
                    isActive: $isSettingsActive
                ) {
                    Button(action: { isSettingsActive.toggle() }) {
                        Image(systemName: "gear")
                    }
                }
            )
        }
        .navigationViewStyle(StackNavigationViewStyle())
    }

    var heart: some View {
        Image(systemName: isMenuServing ? "heart.fill" : "heart.slash.fill")
            .foregroundColor(.pink)
            // .scaleEffect(isMenuServing ? 1.1 : 0)
            .animation(Animation.interactiveSpring().delay(0.2))
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView(
            menuViewModel: MenuViewModel(),
            orderViewModel: OrderViewModel(),
            healthViewModel: HealthViewModel()
        )
    }
}
