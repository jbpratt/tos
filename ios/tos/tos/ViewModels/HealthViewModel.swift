import GRPC
import NIO
import SwiftUI

final class HealthViewModel: ChannelViewModel, ObservableObject, Identifiable {
    public typealias Status = Grpc_Health_V1_HealthCheckResponse.ServingStatus

    @Published private(set) var services: [String: Status] = [:]
    private(set) var menuServiceName = "tospb.MenuService"
    private(set) var orderServiceName = "tospb.OrderService"
    private lazy var client = Grpc_Health_V1_HealthClient(
        channel: self.conn,
        defaultCallOptions: CallOptions(timeLimit: .timeout(TimeAmount.seconds(2)))
    )

    override init() {
        super.init()

        for service in [menuServiceName, orderServiceName] {
            check(service)
        }

        _ = Timer.scheduledTimer(withTimeInterval: 5, repeats: true) { _ in
            for service in [self.menuServiceName, self.orderServiceName] {
                self.check(service)
            }
        }
    }

    func check(_ service: String) {
        client.check(Grpc_Health_V1_HealthCheckRequest.with { $0.service = service }).response.whenComplete { res in
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
