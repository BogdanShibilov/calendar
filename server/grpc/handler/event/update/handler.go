package update

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/model/event"
	"hwCalendar/protobuf/eventpb"
	"hwCalendar/storage"
)

func Handle(_ context.Context, req *eventpb.UpdateEventRequest) (*emptypb.Empty, error) {
	updatedEvent := eventFromReq(req)

	err := updatedEvent.Update()
	if err != nil {
		return nil, handleError(err)
	}

	return &emptypb.Empty{}, nil
}

func eventFromReq(req *eventpb.UpdateEventRequest) *event.Event {
	return &event.Event{
		Id:          int(req.Event.Id),
		Name:        req.Event.Name,
		Description: req.Event.Description,
		Timestamp:   req.Event.Timestamp.AsTime(),
	}
}

func handleError(err error) error {
	if errors.Is(err, event.ErrInvalidId) ||
		errors.Is(err, event.ErrEmptyName) ||
		errors.Is(err, event.ErrEmptyDescription) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, storage.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
