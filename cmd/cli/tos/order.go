package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/jbpratt/tos/pkg/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var name string

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Place an order",
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}

		defer cc.Close()

		res, err := doSubmitOrderRequest(pb.NewOrderServiceClient(cc))
		if err != nil {
			log.Fatalf("Failed to submit order: %v", err)
		}
		fmt.Println(res)
	},
}

func doSubmitOrderRequest(c pb.OrderServiceClient) (string, error) {
	go fmt.Println("Starting order request")
	req := &pb.Order{
		Name: name,
		Items: []*pb.Item{
			{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Options: []*pb.Option{
				{Name: "pickles", Price: 0, Selected: true, Id: 1},
			}},
		},
		Total: 495,
	}

	res, err := c.SubmitOrder(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.GetResponse(), nil
}

func init() {
	rootCmd.AddCommand(orderCmd)
	orderCmd.Flags().StringVarP(&name, "order name", "n", "Majora", "Name to place order under")
}
