package byusername

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hwCalendar/proto/userpb"
	"hwCalendar/user/model/user"
	"hwCalendar/user/storage"
)

func Handle(ctx context.Context, req *userpb.UserByUsernameRequest) (*userpb.UserByUsernameResponse, error) {
	u, err := user.ByUsername(ctx, req.Username)
	if err != nil {
		return nil, handleError(err)
	}

	return resFromUser(u), nil
}

func resFromUser(u *user.User) *userpb.UserByUsernameResponse {
	return &userpb.UserByUsernameResponse{
		User: &userpb.User{
			Id:        int32(u.Id),
			Username:  u.Username,
			PassHash:  u.PassHash,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
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
