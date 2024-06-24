package add

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/calendar/model/event"
	"hwCalendar/calendar/storage"
	"hwCalendar/proto/eventpb"
)

func Handle(ctx context.Context, req *eventpb.AddEventRequest) (*eventpb.AddEventResponse, error) {
	newEvent := event.New(
		req.Name,
		req.Description,
		req.Timestamp.AsTime(),
	)

	id, err := newEvent.Add(ctx)
	if err != nil {
		return nil, handleError(err)
	}

	return &eventpb.AddEventResponse{
		Id: int32(id),
	}, nil
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
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}

	return status.Errorf(codes.Internal, "add event failed: %v", err)
}
