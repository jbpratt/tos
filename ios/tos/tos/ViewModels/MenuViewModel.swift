import Combine

final class MenuViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private var client: Tospb_MenuServiceClient?
    @Published private(set) var menu: Tospb_Menu? = nil

    override init() {
        super.init()
        client = Tospb_MenuServiceClient(channel: super.channel)
    }

    func getMenu() -> Tospb_Menu {
        #if DEBUG
            return loadMenu()
        #else
            let request = Tospb_Empty()
            let call = client.getMenu(request)
            let response = try? call.response.wait()
            return response!
        #endif
    }

    func createItem(_ item: Tospb_Item) {
        do {
            _ = try client!.createMenuItem(item).response.wait()
        } catch {
            print("createMenuItem failed: \(error)")
            return
        }
    }

    func deleteItem(_ itemID: Int64) {
        let req: Tospb_DeleteMenuItemRequest = .with {
            $0.id = itemID
        }
        do {
            _ = try client!.deleteMenuItem(req).response.wait()
        } catch {
            print("deleteMenuItem failed: \(error)")
        }
    }

    func updateItem(_ item: Tospb_Item) {
        do {
            _ = try client!.updateMenuItem(item).response.wait()
        } catch {
            print("updateMenuItem failed: \(error)")
        }
    }
}
