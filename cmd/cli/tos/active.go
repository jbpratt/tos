package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/jbpratt/tos/internal/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// activeCmd represents the active command
var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "get all active orders",
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}
		defer cc.Close()

		res, err := requestOrders(pb.NewOrderServiceClient(cc))
		if err != nil {
			log.Fatalf("Failed to submit order: %v", err)
		}
		fmt.Println(res)
	},
}

func requestOrders(c pb.OrderServiceClient) ([]*pb.Order, error) {
	res, err := c.ActiveOrders(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}

	return res.GetOrders(), nil
}

func init() {
	rootCmd.AddCommand(activeCmd)
}
