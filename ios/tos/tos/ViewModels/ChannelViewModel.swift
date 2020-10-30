import Combine
import GRPC
import Logging
import NIO
import NIOTransportServices

class ChannelViewModel {
    private let group = NIOTSEventLoopGroup()
    let conn: ClientConnection
    let logger: Logger = Logger(label: "TOS")

    init() {
        var conf = ClientConnection.Configuration(
            target: .hostAndPort("localhost", 50051),
            eventLoopGroup: group,
            connectionBackoff: nil,
            connectionKeepalive: ClientConnectionKeepalive(
                interval: .seconds(5),
                timeout: .seconds(3)
            ),
            backgroundActivityLogger: self.logger
        )
        let delegate = RecordingDelegate()
        conf.errorDelegate = delegate
        conf.connectivityStateDelegate = delegate
        conn = ClientConnection(configuration: conf)
    }

    deinit {
        let _ = conn.close()
        try! group.syncShutdownGracefully()
    }
}

class RecordingDelegate: ConnectivityStateDelegate, ClientErrorDelegate {
    var errors: [Error] = []
    var connectivity = CurrentValueSubject<ConnectivityState, Never>(.idle)
    
    func connectivityStateDidChange(from oldState: ConnectivityState, to newState: ConnectivityState) {
        connectivity.value = newState
    }
    
    func didCatchError(_ error: Error, logger: Logger, file: StaticString, line: Int) {
        print(error)
        errors.append(error)
    }
}
