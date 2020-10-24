import Foundation

extension Tospb_Order {
    func totalPrice() -> Float {
        self.tax() + self.subTotal()
    }
    
    func subTotal() -> Float {
        items.reduce(0) { x, y in x + y.totalPrice() }
    }
    
    func tax() -> Float {
        0.04 * self.subTotal()
    }
}

extension Tospb_Item {
    func totalPrice() -> Float {
        price + options.filter { $0.selected }.reduce(0) { x, y in x + y.price }
    }
}
