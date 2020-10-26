import NotificationBannerSwift
import SwiftUI


struct EditItemView: View {
    @Binding var item: Tospb_Item

    var body: some View {
        HStack {
            TextField("Item name", text: $item.name)
            ForEach(item.options.indices) { idx in
                HStack {
                    TextField(
                        "Option name",
                        text: $item.options[idx].name,
                        onEditingChanged: { _ in print("changed") },
                        onCommit: { print("commit") }
                    )
                    Toggle(isOn: $item.options[idx].selected) {}
                }
            }
        }
    }
}

struct OptionsListView: View {
    @Binding var item: Tospb_Item

    var body: some View {
        Text(item.name).font(.headline)
        ForEach(item.options, id: \.self) { opt in
            HStack {
                Text(opt.name)
                Spacer()
                PriceView(price: opt.price)
                if opt.selected {
                    Image(systemName: "checkmark")
                }
            }.onTapGesture {
                if let idx = item.options.firstIndex(of: opt) {
                    item.options[idx].selected = !item.options[idx].selected
                }
            }
        }
    }
}

struct BottomBarView: View {
    var price: Float
    var onSubmit: () -> Void
    var onCancel: () -> Void

    init(_ price: Float, onSubmit: @escaping () -> Void, onCancel: @escaping () -> Void) {
        self.price = price
        self.onSubmit = onSubmit
        self.onCancel = onCancel
    }

    var body: some View {
        HStack {
            Button(action: onSubmit) {
                Image(systemName: "plus.circle")
            }
            Spacer()
            PriceView(price: price)
            Spacer()
            Button(action: onCancel) {
                Image(systemName: "xmark.circle")
            }
        }
        .padding()
    }
}

struct PopupMenuView: View {
    @ObservedObject var viewModel: OrderViewModel
    @Binding var item: Tospb_Item
    @Binding var isActive: Bool
    @State private var editing = false
    @State private var editedItem = Tospb_Item()

    var body: some View {
        VStack {
            if editing {
                EditItemView(item: $item)
            } else {
                OptionsListView(item: $item)
            }
            Divider()
            // Rather than rendering two bars, just conditionally change the action
            BottomBarView(item.totalPrice(), onSubmit: {
                if !editing {
                    viewModel.addToOrder(item)
                    // item = nil
                    StatusBarNotificationBanner(title: "\(item.name) has been added to the order.", style: .success).show()
                } else {
                    // save
                    editing = false
                    item = editedItem
                }
            }, onCancel: {
                if editing {
                    editing = false
                } else {
                    // item = nil
                    isActive = false
                }
            })
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
            item: Binding.constant(loadMenu().categories[0].items[0]),
            isActive: Binding.constant(false)
        )
        .previewLayout(.sizeThatFits)
    }
}
