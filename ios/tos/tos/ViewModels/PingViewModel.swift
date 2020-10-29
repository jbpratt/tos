import Foundation
import NIO
import GRPC

class PingViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private lazy var pingClient = Tospb_PingServiceClient(channel: self.client)
    private var timer: Timer?

    @Published var active: Bool = false

    override init() {
        super.init()
        self.ping()
        timer = Timer.scheduledTimer(withTimeInterval: 60, repeats: true) { timer in
            self.ping()
        }
    }

    func ping() {
        pingClient.sendPing(Tospb_Ping.with {
            $0.message = "ping"
        }).response.whenComplete { res in
            DispatchQueue.main.async {
                switch res {
                case .success(_):
                    self.active = true
                case .failure(let err):
                    self.active = false
                    print("sendPing failed: \(err)")
                }
            }
        }
    }
}
