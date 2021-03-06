package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/jbpratt/tos/internal/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var orderID int

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete an order by supplying the order ID",
	Run: func(cmd *cobra.Command, args []string) {
		if orderID == 0 {
			log.Fatal("must supply an order ID")
		}
		cc, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}

		defer cc.Close()

		res, err := completeOrder(pb.NewOrderServiceClient(cc))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	},
}

func completeOrder(c pb.OrderServiceClient) (string, error) {
	req := &pb.CompleteOrderRequest{
		Id: int64(orderID),
	}
	res, err := c.CompleteOrder(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.GetResponse(), nil
}

func init() {
	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().IntVarP(&orderID, "order identifier", "o", 0, "ID of order to complete")
}
