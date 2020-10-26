import SwiftUI

struct CategoryView: View {
    var category: Tospb_Category

    // I'm not exactly sure yet, but since `isExpanded` is a const,
    // we may be forcing the redraw to occur on all categories in
    // MenuView rather than the one changed. Maybe moving it to be a
    // binding solves this?
    let isExpanded: Bool
    @Binding var itemSelected: Tospb_Item?

    var body: some View {
        VStack {
            Text(category.name).font(.headline)
            if isExpanded {
                ForEach(category.items, id: \.self) { item in
                    ItemView(item: item)
                        .onTapGesture {
                            itemSelected = item
                            print(item.name)
                        }
                }
            }
        }
        .contentShape(Rectangle())
        .frame(minWidth: 0, maxWidth: .infinity)
    }
}
