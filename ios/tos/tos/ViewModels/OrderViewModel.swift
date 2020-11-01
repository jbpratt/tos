import GRPC
import NIO
import SwiftUI

final class OrderViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private lazy var client = Tospb_OrderServiceClient(channel: self.conn)
    @Published private(set) var activeOrders: [Tospb_Order] = []
    @Published private(set) var currentOrder = Tospb_Order()
    @Published var currentOrderName = "" {
        didSet {
            if currentOrderName.isEmpty {
                currentOrder.name = currentOrderName
            }
        }
    }

    override init() {
        super.init()
    }

    func submitOrder() {
        currentOrder.total = currentOrder.totalPrice()
        currentOrder.name = currentOrderName
        guard !currentOrder.name.isEmpty else {
            logger.error("currentOrder.name is empty")
            return
        }
        guard currentOrder.items.count > 0 else {
            logger.error("currentOrder.items is empty")
            return
        }

        client.submitOrder(
            currentOrder
            //callOptions: CallOptions(timeLimit: .timeout(TimeAmount.seconds(5)))
        ).response.whenComplete { res in
            DispatchQueue.main.async {
                switch res {
                case .success(let res):
                    self.logger.info("submitOrder success: \(res)")
                case .failure(let err):
                    self.logger.info("submitOrder failed: \(err)")
                }
            }
        }

        currentOrder = Tospb_Order()
        currentOrderName = ""
    }

    func addToOrder(_ item: Tospb_Item) {
        logger.info("added \(item.name) to order")
        currentOrder.items.append(item)
    }

    func editItemInOrder(_ item: Tospb_Item, _ newItem: Tospb_Item) {
        if let idx = currentOrder.items.firstIndex(of: item) {
            currentOrder.items[idx] = newItem
            logger.info("edited \(idx) in order items")
        }
    }

    func removeFromOrder(_ item: Tospb_Item) {
        if let idx = currentOrder.items.firstIndex(of: item) {
            currentOrder.items.remove(at: idx)
            logger.info("removed \(idx) from order items")
        }
    }

    func listActiveOrders() {
        client.activeOrders(Tospb_Empty()).response.whenComplete { res in
            DispatchQueue.main.async {
                switch res {
                case .success(let res):
                    self.logger.info("activeOrders success: \(res)")
                    self.activeOrders = res.orders
                case .failure(let err):
                    self.logger.info("activeOrders failed: \(err)")
                }
            }
        }
    }
    
    func cancelOrder(_ order: Tospb_Order) {
        logger.info("cancelOrder is unimplemented")
    }
}
