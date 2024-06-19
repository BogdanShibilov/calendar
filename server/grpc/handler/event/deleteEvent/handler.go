package deleteEvent

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/model/event"
	"hwCalendar/protobuf/eventpb"
)

func Handle(_ context.Context, req *eventpb.DeleteEventRequest) (*emptypb.Empty, error) {
	deletedEvent := &event.Event{Id: int(req.Id)}

	err := deletedEvent.Delete()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
