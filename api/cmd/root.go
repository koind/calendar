package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "Microservice calendar",
}

func init() {
	rootCmd.AddCommand(GrpcServerCmd)
	rootCmd.AddCommand(GrpcClientCmd)
	rootCmd.AddCommand(HttpServerCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
