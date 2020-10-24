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
            Spacer()
            Button(action: {
                viewModel.removeFromOrder(item)
                StatusBarNotificationBanner(title: "\(item.name) has been removed from the order.", style: .warning).show()
            }) {
                Image(systemName: "minus")
            }
        }
    }
}
