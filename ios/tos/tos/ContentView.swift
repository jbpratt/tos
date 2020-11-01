import SwiftUI

struct ContentView: View {
    @ObservedObject var menuViewModel: MenuViewModel
    @ObservedObject var orderViewModel: OrderViewModel
    @ObservedObject var healthViewModel: HealthViewModel

    @State private var isActiveOrdersActive: Bool = false
    @State private var isSettingsActive: Bool = false

    var isMenuServing: Bool {
        // Is there really no one that exports the service name?
        healthViewModel.serviceStatus(healthViewModel.menuServiceName) == .serving
    }

    var body: some View {
        NavigationView {
            GeometryReader { geo in
                HStack {
                    MenuOrderView(vm: orderViewModel)
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
                trailing: HStack {
                    NavigationLink(
                        destination: ActiveOrdersView(vm: orderViewModel),
                        isActive: $isActiveOrdersActive
                    ) {
                        Button(action: { isActiveOrdersActive.toggle() }) {
                            Image(systemName: "plus")
                        }
                    }
                    NavigationLink(
                        destination: SettingsView(viewModel: menuViewModel),
                        isActive: $isSettingsActive
                    ) {
                        Button(action: { isSettingsActive.toggle() }) {
                            Image(systemName: "gear")
                        }
                    }
                }
            )
        }
        .navigationViewStyle(StackNavigationViewStyle())
    }

    var heart: some View {
        ZStack {
            Image(systemName: isMenuServing ? "heart.fill" : "heart.slash.fill")
                .foregroundColor(.pink)
                .animation(Animation.interactiveSpring().delay(2.0))
            Circle()
                .strokeBorder(lineWidth: isMenuServing ? 1 : 35 / 2, antialiased: false)
                .opacity(isMenuServing ? 0 : 1)
                .frame(width: 35, height: 35)
                .foregroundColor(.pink)
                .scaleEffect(isMenuServing ? 1 : 0)
                .animation(isMenuServing ? Animation.easeOut(duration: 1.5)
                    .repeatForever(autoreverses: false) : Animation.default)
        }
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
