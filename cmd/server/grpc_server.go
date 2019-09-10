package main

import (
	"context"
	"github.com/koind/calendar/app/adapter"
	"github.com/koind/calendar/app/config"
	domain "github.com/koind/calendar/app/domain/service"
	"github.com/koind/calendar/app/storage/postgres"
	"github.com/koind/calendar/app/transport/grpc/pb"
	"github.com/koind/calendar/app/transport/grpc/service"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

	lis, err := net.Listen("tcp", cfg.GRPCServer.GetDomain())
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	db, err := adapter.IntPostgres(cfg.Postgres)
	if err != nil {
		log.Fatalf("failing to connect to the database %v", err)
	}

	eventRepository := postgres.NewEventRepository(context.Background(), db, *logger)
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
