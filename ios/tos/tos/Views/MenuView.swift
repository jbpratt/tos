import SwiftUI

struct MenuView: View {
    @ObservedObject var menuViewModel: MenuViewModel
    @ObservedObject var orderViewModel: OrderViewModel

    @State private var selection: Set<Tospb_Category> = []
    @State private var itemSelected: Tospb_Item?

    var body: some View {
        ZStack {
            ScrollView {
                VStack(alignment: .leading) {
                    ForEach(menuViewModel.menu!.categories, id: \.self) { cat in
                        CategoryView(category: cat, isExpanded: selection.contains(cat), itemSelected: $itemSelected)
                            .padding()
                            .overlay(RoundedRectangle(cornerRadius: 16)
                                .stroke(Color.black, lineWidth: 2))
                            .onTapGesture { selectDeselect(cat) }
                            .animation(.linear(duration: 0.2))
                    }.padding()
                }
            }

            if itemSelected != nil {
                PopupMenu(viewModel: orderViewModel, item: $itemSelected)
                    .padding(.horizontal)
            }
        }
    }

    func selectDeselect(_ category: Tospb_Category) {
        if selection.contains(category) {
            selection.remove(category)
            itemSelected = nil
        } else {
            selection.insert(category)
        }
    }
}
