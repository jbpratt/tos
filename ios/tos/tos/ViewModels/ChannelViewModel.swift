import Combine
import GRPC
import Logging
import NIO
import NIOTransportServices

class ChannelViewModel {
    private let group = NIOTSEventLoopGroup()
    let conn: ClientConnection
    let logger = Logger(label: "TOS")

    init() {
        let delegate = RecordingDelegate()
        let conf = ClientConnection.Configuration(
            target: .hostAndPort("localhost", 50051),
            eventLoopGroup: group,
            errorDelegate: delegate,
            connectivityStateDelegate: delegate,
            // connectionKeepalive: ClientConnectionKeepalive(
            //    interval: .seconds(10),
            //    timeout: .seconds(3)
            // ),
            backgroundActivityLogger: logger
        )
        conn = ClientConnection(configuration: conf)
    }

    deinit {
        _ = conn.close()
        try! group.syncShutdownGracefully()
    }
}

class RecordingDelegate: ConnectivityStateDelegate, ClientErrorDelegate {
    var errors: [Error] = []
    var connectivity = CurrentValueSubject<ConnectivityState, Never>(.idle)

    func connectivityStateDidChange(from _: ConnectivityState, to newState: ConnectivityState) {
        connectivity.value = newState
    }

    func didCatchError(_ error: Error, logger _: Logger, file _: StaticString, line _: Int) {
        print(error)
        errors.append(error)
    }
}
