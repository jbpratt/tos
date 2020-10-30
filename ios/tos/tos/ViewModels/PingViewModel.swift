import SwiftUI
import NIO
import GRPC

class PingViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private lazy var client = Tospb_PingServiceClient(channel: self.conn)

    @Published var active: Bool = false

    override init() {
        super.init()
        _ = Timer.scheduledTimer(withTimeInterval: 60, repeats: true) { _ in
            self.ping()
        }
    }
    
    func ping() {
        /*
        let opts = CallOptions()
        let req = Tospb_Ping.with {
            $0.message = "ping"
        }
        let call = pingClient.sendPing(req, callOptions: opts)
         */
        //DispatchQueue.main.async {
        self.client.sendPing(Tospb_Ping.with {
            $0.message = "ping"
        }).response.whenComplete { res in
            print(res)
            switch res {
            case .success(_):
                self.active = true
            case .failure(let err):
                self.active = false
                print("sendPing failed: \(err)")
            }
        }
        //}
    }
}
