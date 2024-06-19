package add

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/model/event"
	"hwCalendar/protobuf/eventpb"
	"hwCalendar/storage"
)

func Handle(_ context.Context, req *eventpb.AddEventRequest) (*eventpb.AddEventResponse, error) {
	newEvent := eventFromReq(req)

	id, err := newEvent.Add()
	if err != nil {
		return nil, handleError(err)
	}

	return &eventpb.AddEventResponse{Id: int32(id)}, nil
}

func eventFromReq(req *eventpb.AddEventRequest) *event.Event {
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
	if errors.Is(err, storage.ErrAlreadyExists) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	if errors.Is(err, storage.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Errorf(codes.Internal, "add event failed: %v", err)
}
