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

func Handle(ctx context.Context, req *eventpb.UpdateEventRequest) (*emptypb.Empty, error) {
	id := int(req.Id)
	e, err := event.ById(ctx, id)
	if err != nil {
		return nil, handleError(err)
	}

	err = e.Update(ctx, req.Name, req.Description)
	if err != nil {
		return nil, handleError(err)
	}

	return &emptypb.Empty{}, nil
}

func handleError(err error) error {
	if errors.Is(err, storage.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, event.ErrInvalidId) ||
		errors.Is(err, event.ErrEmptyName) ||
		errors.Is(err, event.ErrEmptyDescription) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
