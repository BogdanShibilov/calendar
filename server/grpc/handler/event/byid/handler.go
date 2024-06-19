package byid

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hwCalendar/model/event"
	"hwCalendar/protobuf/eventpb"
	"hwCalendar/storage"
)

func Handle(_ context.Context, req *eventpb.EventByIdRequest) (*eventpb.EventByIdResponse, error) {
	id := int(req.Id)

	e, err := event.ById(id)
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

	return status.Error(codes.Internal, err.Error())
}
