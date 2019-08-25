package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/koind/calendar/app/config"
	httpServer "github.com/koind/calendar/app/delivery/http/server"
	httpService "github.com/koind/calendar/app/delivery/http/service"
	"github.com/koind/calendar/app/domain/repository"
	"github.com/koind/calendar/app/domain/service"
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	calendarDomain := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	eventService := service.NewEventService(repository.NewDummyEventRepository())
	httpEventService := httpService.NewEventService(eventService, logger)
	hs := httpServer.NewHTTPServer(httpEventService, calendarDomain)

	logger.Error("Error starting app", zap.Error(hs.Start()))
}
