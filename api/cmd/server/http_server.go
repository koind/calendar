package main

import (
	"github.com/koind/calendar/app/config"
	"github.com/koind/calendar/app/domain/service"
	"github.com/koind/calendar/app/storage/memory"
	httpServer "github.com/koind/calendar/app/transport/http/server"
	httpService "github.com/koind/calendar/app/transport/http/service"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
	"log"
)

func main() {
	flag.StringVarP(
		&config.Path,
		"config",
		"c",
		"config/development/config.toml",
		"Путь до конфигурационного toml файла",
	)

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
}
