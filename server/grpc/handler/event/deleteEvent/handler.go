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
	id := int(req.Id)

	err := event.Delete(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
