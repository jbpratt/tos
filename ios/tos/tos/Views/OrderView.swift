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
            Button(action: {}) {
                Text("Submit")
            }.padding(.bottom, 20)
        }
    }
}
