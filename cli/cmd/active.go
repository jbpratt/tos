package cmd

import (
	"context"
	"fmt"
	"log"

	tospb "github.com/jbpratt78/tos/protofiles"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// activeCmd represents the active command
var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "A brief description of your command",
	Long:  `A  quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}
		defer cc.Close()

		res, err := requestOrders(tospb.NewOrderServiceClient(cc))
		if err != nil {
			log.Fatalf("Failed to submit order: %v", err)
		}
		fmt.Println(res)
	},
}

func requestOrders(c tospb.OrderServiceClient) ([]*tospb.Order, error) {
	res, err := c.ActiveOrders(context.Background(), &tospb.Empty{})
	if err != nil {
		return nil, err
	}

	return res.GetOrders(), nil
}

func init() {
	rootCmd.AddCommand(activeCmd)
}
