package cmd

import (
	client "github.com/koind/calendar/api/cmd/client"
	"github.com/koind/calendar/api/cmd/server"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "Microservice calendar",
}

func init() {
	rootCmd.AddCommand(server.GrpcServerCmd)
	rootCmd.AddCommand(server.HttpServerCmd)
	rootCmd.AddCommand(client.GrpcClientCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
