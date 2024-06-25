package update

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/calendar/model/event"
	"hwCalendar/calendar/server/grpc/handler/event/common"
	"hwCalendar/calendar/storage"
	"hwCalendar/proto/eventpb"
)

func Handle(ctx context.Context, req *eventpb.UpdateEventRequest) (*emptypb.Empty, error) {
	err := validate(req)
	if err != nil {
		return nil, err
	}

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

func validate(req *eventpb.UpdateEventRequest) error {
	if req.Name == "" {
		return status.Error(codes.InvalidArgument, common.ErrEmptyName.Error())
	}
	if req.Description == "" {
		return status.Error(codes.InvalidArgument, common.ErrEmptyDescription.Error())
	}

	return nil
}

func handleError(err error) error {
	if errors.Is(err, storage.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
