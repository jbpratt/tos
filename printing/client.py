import sys
import os

import grpc
import argparse
from escpos import printer 

import mookies_pb2
import mookies_pb2_grpc


def run():
    parser = argparse.ArgumentParser()
    parser.add_argument("address", help="server address to dial")
    args = parser.parse_args()
    with grpc.insecure_channel(args.address) as channel:
        order_stub = mookies_pb2_grpc.OrderServiceStub(channel)
        Epson = printer.File("/dev/usb/lp0")
        try:
            print("Dialed " + args.address) 
            for order in order_stub.SubscribeToCompleteOrders(
                    mookies_pb2.SubscribeToOrderRequest(request='pls')):
                Epson.control("LF")
                Epson.set(font='a', height=4, align='center')
                Epson.text("Mookie's Smokehouse'\n")
                Epson.set(font='a', height=2, align='center')
                Epson.text(order.name + "\n")
                Epson.cut()
                print(order)

        except grpc.RpcError as e:
            print("Error raised: " + e.__cause__)

if __name__ == '__main__':
    try:
        run()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
