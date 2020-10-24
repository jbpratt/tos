import SwiftUI

struct OrderView: View {
    @ObservedObject var viewModel: OrderViewModel

    var body: some View {
        VStack {
            Text("Order")
            if viewModel.currentOrder != nil {
                ForEach(viewModel.currentOrder!.items, id: \.self) { item in
                    OrderItemView(viewModel: viewModel, item: item)
                }
                .padding(20)
            }
            Spacer()
            HStack {
                TextField("Enter an order name", text: $viewModel.currentOrderName)
                Button(action: {
                    viewModel.submitOrder()
                }) {
                    Text("Submit")
                }
                .padding(.bottom, 20)
                .disabled(viewModel.currentOrderName.isEmpty)
            }
        }
    }
}
