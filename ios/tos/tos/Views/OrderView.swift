import SwiftUI

struct OrderView: View {
    @ObservedObject var viewModel: OrderViewModel
    
    var body: some View {
        VStack {
            if viewModel.currentOrder != nil {
                ForEach(viewModel.currentOrder!.items, id: \.self) { item in
                    OrderItemView(viewModel: viewModel, item: item)
                }
                .padding()
            } else {
                Text("no items").font(.subheadline).foregroundColor(Color.gray)
            }
            Spacer()
            VStack(alignment: .leading) {
                VStack {
                    HStack {
                        Text("Subtotal:")
                        Spacer()
                        PriceView(price: viewModel.currentOrder?.subTotal() ?? 0.00)
                    }
                    HStack {
                        Text("Tax:")
                        Spacer()
                        PriceView(price: viewModel.currentOrder?.tax() ?? 0.00)
                    }
                    HStack {
                        Text("Total:")
                        Spacer()
                        PriceView(price: viewModel.currentOrder?.totalPrice() ?? 0.00)
                    }
                }.padding().overlay(RoundedRectangle(cornerRadius: 16).stroke(Color.blue, lineWidth: 2))
                HStack {
                    TextField("Enter an order name", text: $viewModel.currentOrderName)
                    Button(action: {
                        viewModel.submitOrder()
                    }) {
                        Text("Submit")
                    }
                    .disabled(viewModel.currentOrderName.isEmpty)
                }
                .padding(.top, 20)
            }
        }
        .padding()
    }
}
