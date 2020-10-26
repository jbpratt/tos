import NotificationBannerSwift
import SwiftUI

struct OrderItemView: View {
    @ObservedObject var viewModel: OrderViewModel
    var item: Tospb_Item

    var body: some View {
        HStack {
            Button(action: {
                viewModel.removeFromOrder(item)
                StatusBarNotificationBanner(title: "\(item.name) has been removed from the order.", style: .warning).show()
            }) {
                Image(systemName: "xmark.circle")
            }
            VStack(alignment: .leading) {
                Text(item.name)
                ForEach(item.options, id: \.self) { opt in
                    if opt.selected {
                        Text("- \(opt.name)").font(.footnote)
                    }
                }
            }
            Spacer()
            PriceView(price: item.totalPrice())
        }
    }
}
