import Combine

final class OrderViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private var client: Tospb_OrderServiceClient?
    @Published var currentOrder: Tospb_Order? = nil
    @Published var currentOrderName = "" {
        didSet {
            if self.currentOrderName.isEmpty {
                currentOrder?.name = currentOrderName
            }
        }
    }

    override init() {
        super.init()
        client = Tospb_OrderServiceClient(channel: super.channel)
    }

    func submitOrder() {
        if currentOrder != nil {
            currentOrder?.total = (currentOrder?.totalPrice())!
            currentOrder?.name = currentOrderName
        }
        guard let order = currentOrder else {
            print("currentOrder is nil")
            return
        }
        guard !order.name.isEmpty else {
            print("currentOrder.name is empty")
            return
        }
        guard order.items.count > 0 else {
            print("currentOrder.items is empty")
            return
        }

        do {
            _ = try client!.submitOrder(order).response.wait()
        } catch {
            print("submitOrder failed: \(error)")
        }

        currentOrder = nil
        currentOrderName = ""
    }

    func addToOrder(_ item: Tospb_Item) {
        if currentOrder == nil {
            currentOrder = Tospb_Order()
        }

        print("added \(item.name) to order")
        currentOrder?.items.append(item)
    }

    func removeFromOrder(_ item: Tospb_Item) {
        guard currentOrder != nil else {
            return
        }

        if let idx = currentOrder!.items.firstIndex(of: item) {
            currentOrder?.items.remove(at: idx)
            print("removed \(idx) from order items")
        }
    }

    func activeOrders() -> [Tospb_Order] {
        var orders: [Tospb_Order] = []
        do {
            let resp = try client!.activeOrders(Tospb_Empty()).response.wait()
            orders = resp.orders
        } catch {
            print("failed to get active orders: \(error)")
        }
        return orders
    }
}
