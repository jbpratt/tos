import SwiftUI

struct ItemView: View {
    var item: Tospb_Item

    var body: some View {
        HStack {
            Text(item.name)
            Spacer()
            PriceView(price: item.price)
        }
        .padding(.top, 10)
    }
}
