package all

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hwCalendar/calendar/model/event"
	"hwCalendar/proto/eventpb"
)

func Handle(ctx context.Context, req *eventpb.AllEventsRequest) (*eventpb.AllEventsResponse, error) {
	totalPages, err := getTotalPages(ctx, int(req.Limit))
	if err != nil {
		return nil, handleError(err)
	}

	if totalPages < int(req.Page) {
		return nil, status.Error(codes.InvalidArgument, "too many pages")
	}

	all, err := event.All(ctx, int(req.Limit), int(req.Page))
	if err != nil {
		return nil, handleError(err)
	}

	return &eventpb.AllEventsResponse{
		Events:     pbFromEvents(all),
		TotalPages: int32(totalPages),
	}, nil
}

func getTotalPages(ctx context.Context, limit int) (int, error) {
	eventCount, err := event.Count(ctx)
	if err != nil {
		return 0, status.Error(codes.Internal, err.Error())
	}

	totalPages := eventCount / limit
	if eventCount%limit != 0 {
		totalPages++
	}
	return totalPages, nil
}

func pbFromEvents(allEvents []event.Event) []*eventpb.Event {
	eventPbSlice := make([]*eventpb.Event, 0, len(allEvents))
	for _, e := range allEvents {
		eventPbSlice = append(eventPbSlice, &eventpb.Event{
			Id:          int32(e.Id),
			Name:        e.Name,
			Description: e.Description,
			Timestamp:   timestamppb.New(e.Timestamp),
		})
	}

	return eventPbSlice
}

func handleError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
