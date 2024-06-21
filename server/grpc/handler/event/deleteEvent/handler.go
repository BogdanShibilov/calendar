package deleteEvent

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/model/event"
	"hwCalendar/protobuf/eventpb"
)

func Handle(ctx context.Context, req *eventpb.DeleteEventRequest) (*emptypb.Empty, error) {
	id := int(req.Id)

	err := event.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
