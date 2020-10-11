import SwiftUI

struct PriceView: View {
    var price: Float

    var body: some View {
        Text(String(format: "$%.2f", price / 100))
    }
}
