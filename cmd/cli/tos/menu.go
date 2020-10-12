package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/jbpratt/tos/pkg/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "retrieve the entire menu",
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}
		defer cc.Close()

		categories, err := doMenuRequest(pb.NewMenuServiceClient(cc))
		if err != nil {
			log.Fatalf("Failed to get menu: %v", err)
		}
		fmt.Println(categories)
	},
}

func doMenuRequest(c pb.MenuServiceClient) ([]*pb.Category, error) {
	fmt.Println("Starting to request menu...")
	res, err := c.GetMenu(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}
	return res.GetCategories(), nil
}

func init() {
	rootCmd.AddCommand(menuCmd)
}
