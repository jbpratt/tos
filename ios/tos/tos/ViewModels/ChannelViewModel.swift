import NIO
import GRPC

class ChannelViewModel {
    private let group = MultiThreadedEventLoopGroup(numberOfThreads: 1)
    let channel: GRPCChannel

    init() {
        channel = ClientConnection.insecure(group: group).connect(host: "localhost", port: 50051)
    }

    deinit {
        // this maybe should be defer'd
        try? group.syncShutdownGracefully()
    }
}
