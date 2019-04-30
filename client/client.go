package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	mookiespb "github.com/jbpratt78/mookies/protofiles"
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

	c := mookiespb.NewMenuServiceClient(cc)

	defer cc.Close()
	doMenuRequest(c)
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
