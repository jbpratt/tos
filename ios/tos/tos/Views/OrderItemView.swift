import NotificationBannerSwift
import SwiftUI

struct OrderItemView: View {
    @ObservedObject var viewModel: OrderViewModel
    var item: Tospb_Item

    var body: some View {
        HStack {
            Text(item.name)
            Spacer()
            PriceView(price: item.price)
            Button(action: {
                viewModel.removeFromOrder(item)
                StatusBarNotificationBanner(title: "\(item.name) has been removed from the order.", style: .warning).show()
            }) {
                Image(systemName: "minus")
            }
            VStack {
                ForEach(item.options, id: \.self) { opt in
                    if opt.selected {
                        Text("\(opt.name)")
                    }
                }
            }
        }
    }
}
