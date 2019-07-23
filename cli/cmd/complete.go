// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"log"

	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var orderID int

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Long:  "Complete an order by supplying the order ID",
	Run: func(cmd *cobra.Command, args []string) {
		if orderID == 0 {
			log.Fatal("must supply an order ID")
		}
		cc, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}

		orderClient := mookiespb.NewOrderServiceClient(cc)

		defer cc.Close()
		res, err := completeOrder(orderClient)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	},
}

func completeOrder(c mookiespb.OrderServiceClient) (string, error) {
	req := &mookiespb.CompleteOrderRequest{
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
