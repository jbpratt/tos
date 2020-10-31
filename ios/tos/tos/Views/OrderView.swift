import SwiftUI

struct OrderView: View {
    @ObservedObject var vm: OrderViewModel

    var pricebar: some View {
        VStack {
            HStack {
                Text("Subtotal:")
                Spacer()
                PriceView(price: vm.currentOrder.subTotal())
            }
            HStack {
                Text("Tax:")
                Spacer()
                PriceView(price: vm.currentOrder.tax())
            }
            HStack {
                Text("Total:")
                Spacer()
                PriceView(price: vm.currentOrder.totalPrice())
            }
        }
        .padding()
        .overlay(RoundedRectangle(cornerRadius: 16)
                    .stroke(Color.black, lineWidth: 2))
    }

    var body: some View {
        VStack {
            ScrollView {
                if !vm.currentOrder.items.isEmpty {
                    ForEach(vm.currentOrder.items, id: \.self) { item in
                        OrderItemView(viewModel: vm, item: item)
                    }
                    .padding([.top, .bottom], 5)
                } else {
                    Text("no items")
                        .font(.subheadline)
                        .foregroundColor(Color.gray)
                }
            }
            Spacer()
            VStack(alignment: .leading) {
                pricebar
                HStack {
                    TextField("Enter an order name", text: $vm.currentOrderName)
                    Button(action: {
                        vm.submitOrder()
                    }) {
                        Image(systemName: "arrow.right.circle")
                    }
                    .disabled(vm.currentOrderName.isEmpty)
                }
                .padding(.top, 20)
            }
        }
        .padding()
    }
}
