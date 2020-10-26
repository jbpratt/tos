import NotificationBannerSwift
import SwiftUI

struct PopupMenuView: View {
    @ObservedObject var viewModel: OrderViewModel
    @Binding var item: Tospb_Item?
    @State private var editing = false
    @State private var editedItem = Tospb_Item()

    var body: some View {
        VStack {
            Unwrap(item) { i in
                HStack {
                    Spacer()
                    Button(action: {
                        editing = !editing
                        editedItem = i
                    }) {
                        Image(systemName: "gear")
                    }
                    .padding([.top, .trailing])
                }
                if editing {
                    TextField(
                        "Item name",
                        text: $editedItem.name,
                        onEditingChanged: { _ in print("changed") },
                        onCommit: { print("commit") }
                    )
                    ForEach(editedItem.options.indices) { idx in
                        HStack {
                            TextField(
                                "Option name",
                                text: $editedItem.options[idx].name,
                                onEditingChanged: { _ in print("changed") },
                                onCommit: { print("commit") }
                            )
                            Toggle(isOn: $editedItem.options[idx].selected) {}
                        }
                    }
                } else {
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
                        // .padding(.top, 10)
                    }
                }
                Divider()
                // Bottom button bar.
                // Rather than rendering two bars, just conditionally change the action
                HStack {
                    Button(action: {
                        if !editing {
                            viewModel.addToOrder(i)
                            item = nil
                            StatusBarNotificationBanner(title: "\(i.name) has been added to the order.", style: .success).show()
                        } else {
                            // save
                            editing = false
                            item = editedItem
                        }
                    }) {
                        Image(systemName: "plus.circle")
                    }
                    Spacer()
                    PriceView(price: i.totalPrice())
                    Spacer()
                    Button(action: {
                        if editing {
                            editing = false
                        } else {
                            item = nil
                        }
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
            .background(Color.white
                .cornerRadius(16)
                .shadow(radius: 8)))
    }
}

struct PopupMenuView_Preview: PreviewProvider {
    static var previews: some View {
        PopupMenuView(
            viewModel: OrderViewModel(),
            item: Binding.constant(loadMenu().categories[0].items[0])
        )
        .previewLayout(.sizeThatFits)
    }
}
