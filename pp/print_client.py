import grpc

import mookies_pb2
import mookies_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        order_stub = mookies_pb2_grpc.OrderServiceStub(channel)
        menu_stub = mookies_pb2_grpc.MenuServiceStub(channel)

        try:
            for order in order_stub.SubscribeToOrders(mookies_pb2.SubscribeToOrderRequest(request='pls')):
                print(order)
            #response = menu_stub.GetMenu(mookies_pb2.Empty())
            #print(response)
        except grpc.RpcError as e:
            print("Error raised: " + e.details())

if __name__ == '__main__':
    run()
