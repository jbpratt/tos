package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/aarzilli/nucular"
	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", ":50051", "address to dial")
)

func main() {
	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	orderClient := mookiespb.NewOrderServiceClient(cc)

	defer cc.Close()
	subscribeToOrders(orderClient)
}

func subscribeToOrders(c mookiespb.OrderServiceClient) {

	fmt.Println("Subscribing to orders...")
	req := &mookiespb.SubscribeToOrderRequest{
		Request: "send me orders",
	}

	stream, err := c.SubscribeToOrders(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while subscribing to orders RPC: %v", err)
	}
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// end of stream, never hope to hit?
			// or call subscribeToOrders often
			break
		}
		if err != nil {
			log.Fatalf("Error while reading from stream: %v", err)
		}
		log.Printf("Received order: %v", order)
	}

}

func nestedMenu(w *nucular.Window) {
	w.Row(20).Static(180)
	w.Label("Test", "CC")
}
