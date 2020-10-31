import GRPC
import NIO
import SwiftUI

final class HealthViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private(set) var menuServiceName = "tospb.MenuService"
    private(set) var orderServiceName = ""

    public typealias Status = Grpc_Health_V1_HealthCheckResponse.ServingStatus
    private lazy var client = Grpc_Health_V1_HealthClient(channel: self.conn)
    @Published private(set) var services: [String: Status] = [:]

    override init() {
        super.init()

        for service in [menuServiceName] {
            check(service)
        }

        _ = Timer.scheduledTimer(withTimeInterval: 5, repeats: true) { _ in
            for service in [self.menuServiceName] {
                self.check(service)
            }
        }
    }

    func check(_ service: String) {
        client.check(Grpc_Health_V1_HealthCheckRequest.with {
            $0.service = service
        }, callOptions: CallOptions(
            timeLimit: .timeout(TimeAmount.seconds(2))
        )).response.whenComplete { res in
            DispatchQueue.main.async {
                switch res {
                case .success(let res):
                    self.services[service] = res.status
                    self.logger.info("healthCheck: \(service) is \(res.status)")
                case .failure(let err):
                    self.services[service] = .unknown
                    self.logger.error("check failed: \(err)")
                }
            }
        }
    }

    func watch(_ service: String) {}

    func serviceStatus(_ service: String) -> Status {
        guard let status = services[service] else {
            return Status.serviceUnknown
        }
        return status
    }
}
