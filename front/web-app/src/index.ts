import * as grpcWeb from 'grpc-web';

import { MenuServiceClient, ServiceError } from '../generated/mookies_pb_service';
import { Empty, Menu } from '../generated/mookies_pb'; 

class MenuApp {
    static readonly INTERVAL = 500;
    constructor(public menuService: MenuServiceClient) { }

    reqMenu() {
        const request = new Empty();
        this.menuService.getMenu(request,
            (error: ServiceError | null, response: Menu | null) => {
                if (error) {
                    if (error.code !== grpcWeb.StatusCode.OK) {
                        console.log(error)
                    }
                } else {
                    setTimeout(() => {
                        console.log(response);
                    }, MenuApp.INTERVAL);
                }
            });
    }

    load() {
        const self = this;
        this.reqMenu();
    }

}

const menuService = new MenuServiceClient('http://localhost:8080');
const menuApp = new MenuApp(menuService);
menuApp.load();