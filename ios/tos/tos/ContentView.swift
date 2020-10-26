import SwiftUI

struct ContentView: View {
    @ObservedObject var menuViewModel = MenuViewModel()
    @ObservedObject var orderViewModel = OrderViewModel()
    @State var showSettings = false
    @State var itemSelected: Tospb_Item?

    var body: some View {
        NavigationView {
            ZStack {
                GeometryReader { geo in
                    HStack {
                        OrderView(viewModel: orderViewModel, itemSelected: self.$itemSelected)
                        Divider()
                        MenuView(
                            menuViewModel: menuViewModel,
                            orderViewModel: orderViewModel
                        )
                        .frame(minWidth: geo.size.width - (geo.size.width / 3))
                    }
                }
                .navigationBarTitle("Menu", displayMode: .inline)
                .navigationBarItems(trailing: NavigationLink(destination: SettingsView(), isActive: $showSettings) {
                    Button(action: {
                        showSettings = true
                    }) {
                        Text("Settings")
                    }
                })
                
                if itemSelected != nil {
                    PopupMenuView(viewModel: self.orderViewModel, item: self.$itemSelected, editing: Binding.constant(true))
                        .padding()
                }
            }
        }
        .navigationViewStyle(StackNavigationViewStyle())
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
