package deleteUser

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/proto/userpb"
	"hwCalendar/user/model/user"
)

func Handle(ctx context.Context, req *userpb.DeleteUserRequest) (*emptypb.Empty, error) {
	id := int(req.Id)

	err := user.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
