import SwiftUI

struct SettingsView: View {
    @ObservedObject var viewModel: MenuViewModel
    
    var body: some View {
        VStack {
            Unwrap(viewModel.menu) { menu in
                List {
                    ForEach(menu.categories, id: \.self) { cat in
                        Section(header: Text(cat.name)) {
                            ForEach(cat.items, id: \.self) { item in
                                Text(item.name)
                            }
                        }
                    }
                }
            }
        }
    }
}

struct SettingsView_Previews: PreviewProvider {
    static var previews: some View {
        SettingsView(viewModel: MenuViewModel())
            .previewLayout(PreviewLayout.sizeThatFits)
    }
}
