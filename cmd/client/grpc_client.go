package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/golang/protobuf/ptypes"
	"github.com/koind/calendar/app/config"
	"github.com/koind/calendar/app/transport/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	timeout := time.Duration(cfg.ClientTimeout) * time.Millisecond
	calendarDomain := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	clientConn, err := grpc.Dial(calendarDomain, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer clientConn.Close()

	client := pb.NewEventClient(clientConn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
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
			Uuid: response.Uuid,
		},
		Request: &pb.EventRequest{Title: "Buy Audi RS6"},
	}
	response, err = client.Update(ctx, event)
	if err != nil {
		log.Print(err)
	}
	log.Printf("%+v", response)

	eventID := &pb.EventID{
		Uuid: response.Uuid,
	}
	status, err := client.Delete(ctx, eventID)
	if err != nil {
		log.Print(err)
	}
	log.Println(status.Status)
}
