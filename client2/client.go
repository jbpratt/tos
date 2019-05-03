package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/golang/protobuf/ptypes"
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
	//subscribeToOrders(orderClient)
	//err = completeOrder(orderClient)
	err = requestOrders(orderClient)
	if err != nil {
		log.Fatal(err)
	}
}

func completeOrder(c mookiespb.OrderServiceClient) error {
	fmt.Println("Starting complete order request...")
	// take this order req in as param
	req := &mookiespb.Order{
		Id:     1,
		Name:   "Majora",
		Status: false,
		Items: []*mookiespb.Item{
			{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Category: "Sandwich"},
		},
		Total:       495,
		TimeOrdered: ptypes.TimestampNow(),
	}
	res, err := c.CompleteOrder(context.Background(), req)
	if err != nil {
		return err
	}
	log.Printf("Response from CompleteOrder: %v\n", res.GetResult())
	return nil
}

func requestOrders(c mookiespb.OrderServiceClient) error {
	req := &mookiespb.OrdersRequest{
		Request: "send me memes",
	}

	res, err := c.Orders(context.Background(), req)
	if err != nil {
		return err
	}
	log.Printf("Response from Orders: %v\n", res.GetOrders())

	return nil
}

func subscribeToOrders(c mookiespb.OrderServiceClient) error {

	fmt.Println("Subscribing to orders...")
	req := &mookiespb.SubscribeToOrderRequest{
		Request: "send me orders",
	}

	stream, err := c.SubscribeToOrders(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// end of stream, never hope to hit?
			// or call subscribeToOrders often
			break
		}
		if err != nil {
			return nil
		}
		log.Printf("Received order: %v\n", order)
	}
	return nil
}
