import sys
import os

import grpc
import argparse
from escpos import printer 

import mookies_pb2
import mookies_pb2_grpc


GRAPHICS_PATH = os.path.abspath(os.path.join(os.path.dirname( __file__ ), '..', 'assets')) 

def logo():
    image = GRAPHICS_PATH + '/logo-small.png'
    return image

def run(address):
    with grpc.insecure_channel(address) as channel:
        order_stub = mookies_pb2_grpc.OrderServiceStub(channel)
        try: 
            Epson = printer.File('/dev/usb/lp0')
            try:
                for order in order_stub.SubscribeToCompleteOrders(
                        mookies_pb2.SubscribeToOrderRequest(request='Request')):
                    Epson.control('LF')
                    Epson.set(font='a', height=2, align='center')
                    Epson.image(logo())
                    Epson.text('\n')
                    Epson.text(order.name + '\n')
                    Epson.text('Total: $' + str(order.total/100) + '\n')
                    Epson.cut()
                    print('Order #' + str(order.id) + ' has been printed')

            except grpc.RpcError as e:
                print(e)
        except PermissionError as e:
            print(e)

if __name__ == '__main__':
    try:
        parser = argparse.ArgumentParser()
        parser.add_argument(
                '--address', type=str, default='localhost:50051', help='server address to dial')
        args = parser.parse_args()
        run(args.address)
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
