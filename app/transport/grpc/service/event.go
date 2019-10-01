package service

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/koind/calendar/app/domain/repository"
	"github.com/koind/calendar/app/domain/service"
	"github.com/koind/calendar/app/transport/grpc/pb"
)

// GRPC-сервер событий
type EventServer struct {
	EventService service.EventServiceInterface
}

// Создает новое событие
func (s *EventServer) Create(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, errors.New("client cancelled, abandoning.")
	}

	datetime, err := ptypes.Timestamp(req.GetDatetime())
	if err != nil {
		return nil, err
	}

	duration, err := ptypes.Duration(req.GetDuration())
	if err != nil {
		return nil, err
	}

	timeSendNotify, err := ptypes.Timestamp(req.GetDatetime())
	if err != nil {
		return nil, err
	}

	newEvent := repository.Event{
		Title:          req.GetTitle(),
		Datetime:       datetime,
		Duration:       duration,
		Description:    req.GetDescription(),
		UserId:         int(req.GetUserId()),
		TimeSendNotify: timeSendNotify,
	}

	event, err := s.EventService.Create(ctx, newEvent)
	if err != nil {
		return nil, err
	}

	eventResp := &pb.EventResponse{
		Id:             int32(event.ID),
		Title:          event.Title,
		Datetime:       req.GetDatetime(),
		Duration:       req.GetDuration(),
		UserId:         req.GetUserId(),
		TimeSendNotify: req.GetTimeSendNotify(),
	}

	return eventResp, nil
}

// Обновляет событие
func (s *EventServer) Update(ctx context.Context, eventChange *pb.EventChange) (*pb.EventResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, errors.New("client cancelled, abandoning.")
	}

	eventID := int(eventChange.GetEventID().GetId())
	newTitle := eventChange.GetTitle()

	newEvent := repository.Event{
		Title: newTitle,
	}

	event, err := s.EventService.Update(ctx, eventID, newEvent)
	if err != nil {
		return nil, err
	}

	datetime, err := ptypes.TimestampProto(event.Datetime)
	if err != nil {
		return nil, err
	}

	duration := ptypes.DurationProto(event.Duration)

	timeSendNotify, err := ptypes.TimestampProto(event.TimeSendNotify)
	if err != nil {
		return nil, err
	}

	eventResp := &pb.EventResponse{
		Id:             int32(event.ID),
		Title:          event.Title,
		Datetime:       datetime,
		Duration:       duration,
		UserId:         int32(event.ID),
		TimeSendNotify: timeSendNotify,
	}

	return eventResp, nil
}

// Удаляет событие
func (s *EventServer) Delete(ctx context.Context, eventID *pb.EventID) (*pb.EventStatus, error) {
	if ctx.Err() == context.Canceled {
		return nil, errors.New("client cancelled, abandoning.")
	}

	ID := int(eventID.GetId())

	err := s.EventService.Delete(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &pb.EventStatus{Status: "ok"}, nil
}
