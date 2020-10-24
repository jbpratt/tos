import Foundation

extension Tospb_Order {
    func totalPrice() -> Float {
        0.04 * items.reduce(0) { x, y in x + y.totalPrice() }
    }
}

extension Tospb_Item {
    func totalPrice() -> Float {
        price + options.filter { $0.selected }.reduce(0) { x, y in x + y.price }
    }
}
