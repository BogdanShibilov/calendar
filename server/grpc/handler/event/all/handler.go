package all

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hwCalendar/model/event"
	"hwCalendar/protobuf/eventpb"
)

func Handle(_ context.Context, _ *emptypb.Empty) (*eventpb.AllEventsResponse, error) {
	all, err := event.All()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return resFromEvents(all), nil
}

func resFromEvents(allEvents []event.Event) *eventpb.AllEventsResponse {
	eventPbSlice := make([]*eventpb.Event, 0, len(allEvents))
	for _, e := range allEvents {
		eventPbSlice = append(eventPbSlice, &eventpb.Event{
			Id:          int32(e.Id),
			Name:        e.Name,
			Description: e.Description,
			Timestamp:   timestamppb.New(e.Timestamp),
		})
	}

	return &eventpb.AllEventsResponse{
		Events: eventPbSlice,
	}
}
