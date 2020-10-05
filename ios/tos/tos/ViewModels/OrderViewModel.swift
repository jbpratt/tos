import Combine

final class OrderViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private var client: Tospb_OrderServiceClient?
    @Published private(set) var currentOrder: Tospb_Order? = nil

    override init() {
        super.init()
        client = Tospb_OrderServiceClient(channel: super.channel)
    }
}
