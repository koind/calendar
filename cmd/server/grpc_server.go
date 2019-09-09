package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/koind/calendar/app/adapter"
	"github.com/koind/calendar/app/config"
	domain "github.com/koind/calendar/app/domain/service"
	"github.com/koind/calendar/app/storage/postgres"
	"github.com/koind/calendar/app/transport/grpc/pb"
	"github.com/koind/calendar/app/transport/grpc/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

	lis, err := net.Listen("tcp", calendarDomain)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	op := adapter.DBOptions{
		DSN:             cfg.DSN,
		PingTimeout:     430,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: 15,
	}

	db, err := adapter.IntDB(op)
	if err != nil {
		logger.Debug("failing to connect to the database", zap.Error(err))
		log.Fatal("failing to connect to the database")
	}

	eventRepository := postgres.NewEventRepository(db, *logger, context.Background())
	eventService := domain.NewEventService(eventRepository)
	eventServer := &service.EventServer{
		EventService: eventService,
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterEventServer(grpcServer, eventServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
