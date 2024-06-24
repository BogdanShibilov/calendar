package byid

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hwCalendar/calendar/model/event"
	"hwCalendar/calendar/storage"
	"hwCalendar/proto/eventpb"
)

func Handle(ctx context.Context, req *eventpb.EventByIdRequest) (*eventpb.EventByIdResponse, error) {
	id := int(req.Id)

	e, err := event.ById(ctx, id)
	if err != nil {
		return nil, handleError(err)
	}

	return resFromEvent(e), nil
}

func resFromEvent(e *event.Event) *eventpb.EventByIdResponse {
	return &eventpb.EventByIdResponse{
		Event: &eventpb.Event{
			Id:          int32(e.Id),
			Name:        e.Name,
			Description: e.Description,
			Timestamp:   timestamppb.New(e.Timestamp),
		},
	}
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
