import NotificationBannerSwift
import SwiftUI

struct PopupMenuView: View {
    @ObservedObject var viewModel: OrderViewModel
    @Binding var item: Tospb_Item?
    @Binding var editing: Bool

    var body: some View {
        VStack {
            Unwrap(item) { i in
                Text(i.name).font(.headline)
                ForEach(i.options, id: \.self) { opt in
                    HStack {
                        Text(opt.name)
                        Spacer()
                        PriceView(price: opt.price)
                        if opt.selected {
                            Image(systemName: "checkmark")
                        }
                    }.onTapGesture {
                        if let idx = i.options.firstIndex(of: opt) {
                            item?.options[idx].selected = !(item?.options[idx].selected)!
                        }
                    }
                    .padding(.top, 10)
                }
                HStack {
                    Button(action: {
                        if editing {
                            var oldItem: Tospb_Item? = nil
                            viewModel.currentOrder?.items.forEach { orderItem in
                                if item?.orderItemID == orderItem.orderItemID {
                                    oldItem = orderItem
                                }
                            }
                            viewModel.editItemInOrder(oldItem!, i)
                        } else {
                            var item = i
                            if viewModel.currentOrder != nil {
                                print("\(Int64((viewModel.currentOrder?.items.count)!))")
                                item.orderItemID = Int64((viewModel.currentOrder?.items.count)!)
                            } else {
                                item.orderItemID = 0
                            }
                            viewModel.addToOrder(item)
                        }
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
        .background(RoundedRectangle(cornerRadius: 16)
                        .stroke(Color.black, lineWidth: 2)
                        .background((Color.white)
                                        .cornerRadius(16)
                                        .shadow(radius: 8)))
    }
}
