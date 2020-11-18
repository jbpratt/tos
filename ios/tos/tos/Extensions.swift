import Foundation

extension Tospb_Order {
    func totalPrice() -> Float {
        tax() + subTotal()
    }

    func subTotal() -> Float {
        items.reduce(0) { x, y in x + y.totalPrice() }
    }

    func tax() -> Float {
        0.04 * subTotal()
    }
}

extension Tospb_Item {
    func totalPrice() -> Float {
        price + options.filter { $0.selected }.reduce(0) { x, y in x + y.price }
    }

    func copy() -> Self {
        return Tospb_Item.with {
            $0.categoryID = self.categoryID
            $0.id = self.id
            $0.name = self.name
            $0.options = self.options
            $0.orderItemID = self.orderItemID
            $0.price = self.price
        }
    }
}
