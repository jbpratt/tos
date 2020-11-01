import SwiftUI

struct ActiveOrderView: View {
    var order: Tospb_Order
    var cancelOrder: (Tospb_Order) -> Void
    @State private var isExtended: Bool = false
    
    init(_ order: Tospb_Order, cancel: @escaping (Tospb_Order) -> Void) {
        self.order = order
        self.cancelOrder = cancel
    }
    
    var body: some View {
        VStack {
            HStack {
                Text(order.name)
                Spacer()
                Text("$\(order.total)")
                Button(action: {
                    cancelOrder(order)
                }) {
                    Image(systemName: "xmark.circle")
                }
            }.onTapGesture {
                isExtended.toggle()
            }
            
            if isExtended {
                Text("meme")
            }
        }
    }
}

struct ActiveOrdersView: View {
    @ObservedObject var vm: OrderViewModel

    var body: some View {
        VStack {
            Button(action: { vm.listActiveOrders() }) {
                Text("Load orders")
            }
            Spacer()
            ScrollView {
                ForEach(vm.activeOrders, id: \.self) { ord in
                    ActiveOrderView(ord, cancel: vm.cancelOrder)
                }
            }
        }
        .padding()
        .navigationBarTitle("Orders")
    }
}
