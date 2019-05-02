package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aarzilli/nucular"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
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

	//menuClient := mookiespb.NewMenuServiceClient(cc)
	orderClient := mookiespb.NewOrderServiceClient(cc)

	defer cc.Close()
	//doMenuRequest(menuClient)
	doSubmitOrderRequest(orderClient)

	//wnd := nucular.NewMasterWindow(0, "Mookies", nestedMenu)
	//wnd.Main()
}

func doMenuRequest(c mookiespb.MenuServiceClient) {
	fmt.Println("Starting to request menu...")
	req := &empty.Empty{}

	res, err := c.GetMenu(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetMenu RPC: %v\n", err)
	}
	log.Printf("Response from GetMenu: %v\n", res.GetItems())
}

func doSubmitOrderRequest(c mookiespb.OrderServiceClient) {
	fmt.Println("Starting order request")
	req := &mookiespb.SubmitOrderRequest{
		Order: &mookiespb.Order{
			Id:   1,
			Name: "Majora",
			Items: []*mookiespb.Item{
				{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Category: "Sandwich"},
			},
			Total:       495,
			TimeOrdered: &timestamp.Timestamp{},
		},
	}

	res, err := c.SubmitOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while submitting order RPC: %v\n", err)
	}
	log.Printf("Response from SubmitOrder: %v\n", res.GetResult())
}

func nestedMenu(w *nucular.Window) {
	w.Row(20).Static(180)
	w.Label("Test", "CC")
}
