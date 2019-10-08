package cmd

import (
	"github.com/koind/calendar/api/app/config"
	"github.com/koind/calendar/api/app/domain/service"
	"github.com/koind/calendar/api/app/storage/memory"
	httpServer "github.com/koind/calendar/api/app/transport/http/server"
	httpService "github.com/koind/calendar/api/app/transport/http/service"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
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

		eventService := service.NewEventService(memory.NewEventRepository())
		httpEventService := httpService.NewEventService(eventService, logger)
		hs := httpServer.NewHTTPServer(httpEventService, cfg.HTTPServer.GetDomain())

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
