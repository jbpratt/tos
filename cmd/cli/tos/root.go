// Package cmd implements a CLI for debugging
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Address is the server address to dial when making a connection
var Address string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tos",
	Short: "CLI client tool for interacting with the tos GRPC server",
	Long:  "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&Address, "address", "a", ":50051", "Address to dial")
}
