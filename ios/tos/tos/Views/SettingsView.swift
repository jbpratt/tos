import SwiftUI

struct EditItemView: View {
    @Binding var item: Tospb_Item
    
    var body: some View {
        VStack {
            HStack {
                TextField("Name", text: $item.name)
                Spacer()
                TextField("Price", value: $item.price, formatter: NumberFormatter())
                    .keyboardType(.numberPad)
            }
            ForEach(item.options.indices) { idx in
                HStack {
                    TextField("Option name", text: $item.options[idx].name)
                    Spacer()
                    TextField("Price", value: $item.options[idx].price, formatter: NumberFormatter())
                        .keyboardType(.numberPad)
                    Spacer()
                    Toggle(isOn: $item.options[idx].selected){}
                }
            }
        }
        .padding()
    }
}

struct SettingsView: View {
    @ObservedObject var viewModel: MenuViewModel
    @State private var editedItem = Tospb_Item()
    @State private var isSheetActive = false
    
    var body: some View {
        VStack {
            Unwrap(viewModel.menu) { menu in
                List {
                    ForEach(menu.categories, id: \.self) { cat in
                        Section(header: Text(cat.name)) {
                            ForEach(cat.items, id: \.self) { item in
                                Button(action: {
                                    editedItem = item
                                    isSheetActive = !isSheetActive
                                }) {
                                    HStack {
                                        Text(item.name)
                                        Spacer()
                                        PriceView(price: item.price)
                                    }
                                }.sheet(isPresented: $isSheetActive, onDismiss: {
                                   // viewmodel.updateItem(editedItem)
                                }) {
                                    EditItemView(item: $editedItem)
                                }
                                /*
                                .alert(item: $isSheetActive) {
                                    Alert(title: Text("Save the update"), message: Text("Are you sure?"), primaryButton: .destructive(Text("Yes"), action: {
                                    }), secondaryButton: .cancel())
                                })
                                */
                            }
                        }
                    }
                }
                .listStyle(GroupedListStyle())
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
