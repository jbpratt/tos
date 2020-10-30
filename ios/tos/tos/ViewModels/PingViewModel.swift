import SwiftUI
import NIO
import GRPC

class PingViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private lazy var client = Tospb_PingServiceClient(channel: self.conn)

    @Published private(set) var active: Bool = false
    //@Published private(set) var connectivity: ConnectivityState = .idle

    override init() {
        super.init()
        self.ping()
        _ = Timer.scheduledTimer(withTimeInterval: 60, repeats: true) { _ in
            self.ping()
        }
    }
    
    func ping() {
        self.client.sendPing(Tospb_Ping.with {
            $0.message = "ping"
        }).response.whenComplete { res in
            DispatchQueue.main.async {
                switch res {
                case .success(_):
                    self.active = true
                    print("sendPing success: \(res)")
                case .failure(let err):
                    self.active = false
                    print("sendPing failed: \(err)")
                }
            }
        }
    }
}
