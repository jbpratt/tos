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

	"github.com/golang/protobuf/ptypes"
	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Place an order",
	Long:  `Place an order`,
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}

		//menuClient := mookiespb.NewMenuServiceClient(cc)
		orderClient := mookiespb.NewOrderServiceClient(cc)

		defer cc.Close()
		res, err := doSubmitOrderRequest(orderClient)
		if err != nil {
			log.Fatalf("Failed to submit order: %v", err)
		}
		fmt.Println(res)
	},
}

func doSubmitOrderRequest(c mookiespb.OrderServiceClient) (string, error) {
	fmt.Println("Starting order request")
	req := &mookiespb.SubmitOrderRequest{
		Order: &mookiespb.Order{
			Id:   1,
			Name: "Majora",
			Items: []*mookiespb.Item{
				{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Category: "Sandwich"},
			},
			Total:       495,
			TimeOrdered: ptypes.TimestampNow(),
		},
	}

	res, err := c.SubmitOrder(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.GetResult(), nil
}

func init() {
	rootCmd.AddCommand(orderCmd)
}
