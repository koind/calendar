package server

import (
	"context"
	"github.com/koind/calendar/api/app/config"
	"github.com/koind/calendar/api/app/db"
	"github.com/koind/calendar/api/app/domain/service"
	"github.com/koind/calendar/api/app/storage/postgres"
	httpServer "github.com/koind/calendar/api/app/transport/http/server"
	httpService "github.com/koind/calendar/api/app/transport/http/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

var HttpServerCmd = &cobra.Command{
	Use:   "http_server",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Init(config.Path)

		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}
		defer logger.Sync()

		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(cfg.Postgres.PingTimeout)*time.Millisecond,
		)
		defer cancel()

		pg, err := db.IntPostgres(ctx, config.Postgres(cfg.Postgres))
		if err != nil {
			log.Fatalf("failing to connect to the database %v", err)
		}

		eventRepository := postgres.NewEventRepository(context.Background(), pg, *logger)
		eventService := service.NewEventService(eventRepository)
		httpEventService := httpService.NewEventService(eventService, logger)
		hs := httpServer.NewHTTPServer(httpEventService, cfg.HTTPServer.GetDomain())

		go http.ListenAndServe(cfg.Prometheus.GetPort(), promhttp.Handler())

		logger.Error("Error starting app", zap.Error(hs.Start()))
	},
}

func init() {
	HttpServerCmd.Flags().StringVarP(
		&config.Path,
		"config",
		"c",
		"config/development/config.toml",
		"Путь до конфигурационного toml файла",
	)
}
