import GRPC
import NIO

class ChannelViewModel {
    private let group = MultiThreadedEventLoopGroup(numberOfThreads: 2)
    let client: ClientConnection

    init() {
        let conf = ClientConnection.Configuration(
            target: .hostAndPort("localhost", 50051),
            eventLoopGroup: group,
            connectionKeepalive: ClientConnectionKeepalive(
                interval: .seconds(15),
                timeout: .seconds(10)
            )
        )
        client = ClientConnection(configuration: conf)
    }

    deinit {
        let _ = client.close()
        // this maybe should be defer'd
        try! group.syncShutdownGracefully()
    }
}
