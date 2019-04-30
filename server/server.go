package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	mookiespb "github.com/jbpratt78/mookies/protofiles"
	"google.golang.org/grpc"
)

var (
	listen = flag.String("listen", ":50051", "listen address")
)

type server struct{}

func (*server) GetMenu(ctx context.Context, empty *empty.Empty) (*mookiespb.Menu, error) {
	fmt.Println("Menu function was invoked")
	res := &mookiespb.Menu{
		Items: []*mookiespb.Menu_Item{
			{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Category: "Sandwich"},
			{Name: "Regular Smoked Pulled Pork", Id: 2, Price: 395, Category: "Sandwich"},
		},
	}
	return res, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening on %q...\n", *listen)

	s := grpc.NewServer()
	mookiespb.RegisterMenuServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
