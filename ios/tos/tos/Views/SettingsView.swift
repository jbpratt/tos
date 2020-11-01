import SwiftUI

struct EditItemView: View {
    @Binding var item: Tospb_Item
    @Binding var isPresented: Bool
    var onSave: (Tospb_Item) -> Void

    var body: some View {
        VStack {
            HStack {
                TextField("Name", text: $item.name)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                Spacer()
                TextField("Price", value: $item.price, formatter: NumberFormatter())
                    .keyboardType(.numberPad)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
            }
            ForEach(item.options.indices) { idx in editItemOption($item.options[idx]) }
            buttons
        }
        .padding()
    }

    func editItemOption(_ option: Binding<Tospb_Option>) -> some View {
        HStack {
            TextField("Option name", text: option.name)
                .textFieldStyle(RoundedBorderTextFieldStyle())
            Spacer()
            TextField("Price", value: option.price, formatter: NumberFormatter())
                .keyboardType(.numberPad)
                .textFieldStyle(RoundedBorderTextFieldStyle())
            Spacer()
            Toggle(isOn: option.selected) {}
        }
        .padding()
    }

    var buttons: some View {
        HStack {
            Button(action: { onSave(item) }) {
                Text("Save")
            }
            Spacer()
            Button(action: { isPresented.toggle() }) {
                Text("Cancel")
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
                        categorySection(cat)
                    }
                }
                .listStyle(GroupedListStyle())
            }
        }
        .navigationBarTitle("Settings")
    }

    func categorySection(_ cat: Tospb_Category) -> some View {
        Section(header: Text(cat.name)) {
            ForEach(cat.items, id: \.self) { item in
                itemView(item)
                    .sheet(isPresented: $isSheetActive, content: {
                        EditItemView(item: $editedItem, isPresented: $isSheetActive, onSave: viewModel.updateItem)
                    })
            }
        }
    }

    func itemView(_ item: Tospb_Item) -> some View {
        Button(action: {
            editedItem = item
            isSheetActive = !isSheetActive
        }) {
            HStack {
                Text(item.name)
                Spacer()
                PriceView(price: item.price)
            }
            .foregroundColor(Color.black)
        }
    }
}

struct SettingsView_Previews: PreviewProvider {
    static var previews: some View {
        SettingsView(viewModel: MenuViewModel())
            .previewLayout(PreviewLayout.sizeThatFits)
    }
}
