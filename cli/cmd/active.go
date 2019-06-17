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

		orderClient := mookiespb.NewOrderServiceClient(cc)

		defer cc.Close()
		res, err := requestOrders(orderClient)
		if err != nil {
			log.Fatalf("Failed to submit order: %v", err)
		}
		fmt.Println(res)

	},
}

func requestOrders(c mookiespb.OrderServiceClient) ([]*mookiespb.Order, error) {
	req := &mookiespb.Empty{}

	res, err := c.ActiveOrders(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res.GetOrders(), nil
}

func init() {
	rootCmd.AddCommand(activeCmd)
}
