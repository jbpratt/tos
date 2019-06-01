import grpc
from escpos import printer 

import mookies_pb2
import mookies_pb2_grpc

Epson = printer.File("/dev/usb/lp0")

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        order_stub = mookies_pb2_grpc.OrderServiceStub(channel)
    
        try:
            for order in order_stub.SubscribeToCompleteOrders(mookies_pb2.SubscribeToOrderRequest(request='pls')):
                Epson.text(order.__name__)
                print(order)

        except grpc.RpcError as e:
            print("Error raised: " + e.details())

if __name__ == '__main__':
    run()
