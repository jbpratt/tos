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

	"github.com/golang/protobuf/ptypes/empty"
	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "A brief description of your command",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}
		defer cc.Close()

		c := mookiespb.NewMenuServiceClient(cc)
		categories, err := doMenuRequest(c)
		if err != nil {
			log.Fatalf("Failed to get menu: %v", err)
		}
		fmt.Println(categories)
	},
}

func doMenuRequest(c mookiespb.MenuServiceClient) ([]*mookiespb.Category, error) {
	fmt.Println("Starting to request menu...")
	req := &empty.Empty{}

	res, err := c.GetMenu(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.GetCategories(), nil
}

func init() {
	rootCmd.AddCommand(menuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// menuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// menuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
