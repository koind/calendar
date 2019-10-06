package cmd

import (
	"context"
	"github.com/koind/calendar/notifier/internal"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "notifier",
	Short: "Microservice for sending notifications",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := internal.Init(internal.Path)

		conn, err := amqp.Dial(cfg.RabbitMQ.URL)
		if err != nil {
			log.Fatalf("failing to connect to the rabbitmq %v", err)
		}
		defer conn.Close()

		err = internal.Run(context.Background(), *conn)
		if err != nil {
			log.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().StringVarP(
		&internal.Path,
		"config",
		"c",
		"config/development/config.toml",
		"Путь до конфигурационного toml файла",
	)
}
