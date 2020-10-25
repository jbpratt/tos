import SwiftUI

struct ContentView: View {
    @ObservedObject var menuViewModel = MenuViewModel()
    @ObservedObject var orderViewModel = OrderViewModel()

    var body: some View {
        GeometryReader { geo in
            HStack {
                OrderView(viewModel: orderViewModel)
                Divider()
                MenuView(
                    menuViewModel: menuViewModel,
                    orderViewModel: orderViewModel
                )
                .frame(minWidth: geo.size.width - (geo.size.width / 3))
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
