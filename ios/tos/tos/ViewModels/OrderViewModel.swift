import Combine

final class OrderViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private var client: Tospb_OrderServiceClient?
    @Published var currentOrder: Tospb_Order? = nil

    override init() {
        super.init()
        client = Tospb_OrderServiceClient(channel: super.channel)
    }

    func submitOrder() {
        guard let order = currentOrder else {
            print("currentOrder is nil")
            return
        }
        // validate order
        do {
            _ = try client!.submitOrder(order).response.wait()
        } catch {
            print("submitOrder failed: \(error)")
        }
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
}
