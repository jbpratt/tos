package cmd

import (
	"log"

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
	},
}

func init() {
	rootCmd.AddCommand(orderCmd)
	orderCmd.Flags().StringVarP(&name, "order name", "n", "Majora", "Name to place order under")
}
