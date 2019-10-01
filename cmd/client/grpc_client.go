package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/koind/calendar/app/config"
	"github.com/koind/calendar/app/transport/grpc/pb"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"log"
	"time"
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
	clientConn, err := grpc.Dial(cfg.GRPCClient.GetDomain(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer clientConn.Close()

	client := pb.NewEventClient(clientConn)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GRPCClient.GetTimeout())
	defer cancel()

	timestamp, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Println(err)
		return
	}

	notifyTime, err := ptypes.TimestampProto(time.Now().Add(100))
	if err != nil {
		log.Println(err)
		return
	}

	request := &pb.EventRequest{
		Title:          "To buy a car",
		Datetime:       timestamp,
		Duration:       ptypes.DurationProto(100),
		UserId:         123,
		TimeSendNotify: notifyTime,
	}
	response, err := client.Create(ctx, request)
	if err != nil {
		log.Print(err)
	}
	log.Printf("%+v", response)

	event := &pb.EventChange{
		EventID: &pb.EventID{
			Id: response.Id,
		},
		Title: "Buy Audi RS6",
	}
	response, err = client.Update(ctx, event)
	if err != nil {
		log.Print(err)
	}
	log.Printf("%+v", response)

	eventID := &pb.EventID{
		Id: response.Id,
	}
	status, err := client.Delete(ctx, eventID)
	if err != nil {
		log.Print(err)
	}
	log.Println(status.Status)
}
