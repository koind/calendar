package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/koind/calendar/app/adapter"
	"github.com/koind/calendar/app/config"
	"github.com/koind/calendar/app/domain/service"
	"github.com/koind/calendar/app/storage/memory"
	httpServer "github.com/koind/calendar/app/transport/http/server"
	httpService "github.com/koind/calendar/app/transport/http/service"
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

	eventService := service.NewEventService(memory.NewEventRepository())
	httpEventService := httpService.NewEventService(eventService, logger)
	hs := httpServer.NewHTTPServer(httpEventService, calendarDomain)

	logger.Error("Error starting app", zap.Error(hs.Start()))

	dsn := "postgres://root:123123@postgres:5432/events?sslmode=disable"
	op := adapter.DBOptions{
		DSN:             dsn,
		PingTimeout:     430,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: 15,
	}
	db, err := adapter.IntDB(op)
	if err != nil {
		logger.Debug("dont connection", zap.Error(err))
	}
	logger.Info("db", zap.Any("db", db))
}
