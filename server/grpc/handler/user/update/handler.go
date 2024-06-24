package update

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/model/user"
	"hwCalendar/protobuf/userpb"
	"hwCalendar/storage"
)

func Handle(ctx context.Context, req *userpb.UpdateUserRequest) (*emptypb.Empty, error) {
	id := int(req.Id)
	u, err := user.ById(ctx, id)
	if err != nil {
		return nil, handleError(err)
	}

	err = u.Update(ctx, req.Username, req.Password)
	if err != nil {
		return nil, handleError(err)
	}

	return &emptypb.Empty{}, nil
}

func handleError(err error) error {
	if errors.Is(err, storage.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, user.ErrEmptyUsername) ||
		errors.Is(err, user.ErrEmptyPassword) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
