import sys
import os

from prometheus_client import start_http_server
from py_grpc_prometheus.prometheus_client_interceptor import PromClientInterceptor
from escpos import printer 

import grpc
import argparse
import mookies_pb2
import mookies_pb2_grpc


GRAPHICS_PATH = os.path.abspath(os.path.join(os.path.dirname( __file__ ), '..', 'assets')) 
metrics_port = 9002

def logo():
    image = GRAPHICS_PATH + '/logo-small.png'
    return image

def run(address, crt):

    with grpc.intercept_channel(grpc.insecure_channel(address, options=[('grpc.keepalive_time_ms', 10000)]), PromClientInterceptor()) as channel:
        order_stub = mookies_pb2_grpc.OrderServiceStub(channel)
        start_http_server(metrics_port)
        print('connected')
        try: 
            #Epson = printer.File('/dev/usb/lp0')
            try:
                for order in order_stub.SubscribeToCompleteOrders(
                        mookies_pb2.Empty()):
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
        parser.add_argument(
                '--crt', type=str, default='server.crt', help='cert to use when dialing')
        args = parser.parse_args()
        run(args.address, args.crt)
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
