import NotificationBannerSwift
import SwiftUI

struct PopupMenu: View {
    @ObservedObject var viewModel: OrderViewModel
    @Binding var item: Tospb_Item?

    var body: some View {
        VStack {
            Unwrap(item) { i in
                ForEach(i.options, id: \.self) { opt in
                    HStack {
                        Text(opt.name)
                        Spacer()
                        PriceView(price: opt.price)
                        if opt.selected {
                            Image(systemName: "checkmark")
                        }
                    }.onTapGesture {
                        if let idx = item?.options.firstIndex(of: opt) {
                            item?.options[idx].selected = !(item?.options[idx].selected)!
                        }
                    }
                    .padding(.top, 10)
                }
                HStack {
                    Button(action: {
                        viewModel.addToOrder(i)
                        item = nil
                        StatusBarNotificationBanner(title: "\(i.name) has been added to the order.", style: .success).show()
                    }) {
                        Image(systemName: "plus.circle")
                    }
                    Spacer()
                    PriceView(price: i.totalPrice())
                    Spacer()
                    Button(action: {
                        item = nil
                    }) {
                        Image(systemName: "xmark.circle")
                    }
                }
                .padding()
            }
        }
        .padding()
        .background(RoundedRectangle(cornerRadius: 4)
            .stroke(Color.black, lineWidth: 2)
            .background(Color.white))
    }
}
