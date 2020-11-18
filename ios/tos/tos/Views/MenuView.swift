import SwiftUI

struct MenuView: View {
    @ObservedObject var menuViewModel: MenuViewModel
    @ObservedObject var orderViewModel: OrderViewModel

    @State private var isPopupActive: Bool = false
    @State private var selection: Set<Tospb_Category> = []
    @State private var itemSelected = Tospb_Item()

    var body: some View {
        ZStack {
            ScrollView {
                VStack {
                    Unwrap(menuViewModel.menu) { menu in
                        ForEach(menu.categories, id: \.self) { cat in
                            CategoryView(
                                category: cat,
                                isExpanded: selection.contains(cat),
                                itemSelected: $itemSelected,
                                isPopupActive: $isPopupActive
                            )
                            .onTapGesture { selectDeselect(cat) }
                        }
                        .padding()
                    }
                }
            }

            if isPopupActive {
                PopupMenuView(viewModel: orderViewModel, item: $itemSelected, isActive: $isPopupActive)
                    .padding(50)
            }
        }
    }

    func selectDeselect(_ category: Tospb_Category) {
        if selection.contains(category) {
            selection.remove(category)
        } else {
            selection.insert(category)
        }
        isPopupActive = false
    }
}

struct MenuView_Previews: PreviewProvider {
    static var previews: some View {
        MenuView(menuViewModel: MenuViewModel(), orderViewModel: OrderViewModel())
            .previewLayout(PreviewLayout.sizeThatFits)
    }
}
