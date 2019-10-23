package cmd

import (
	"context"
	"github.com/koind/calendar/scheduler/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"sync"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "A microservice that searches for events to be notified",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := internal.Init(internal.Path)

		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(cfg.Postgres.PingTimeout)*time.Second,
		)
		defer cancel()

		db, err := internal.IntPostgres(ctx, cfg.Postgres)
		if err != nil {
			log.Fatalf("failing to connect to the database %v", err)
		}

		conn, err := amqp.Dial(cfg.RabbitMQ.URL)
		if err != nil {
			log.Fatalf("failing to connect to the rabbitmq %v", err)
		}
		defer conn.Close()

		ticker := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ticker.C:
					err = internal.Run(context.Background(), db, *conn)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			http.ListenAndServe(cfg.Prometheus.GetPort(), promhttp.Handler())
		}()

		wg.Wait()
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
