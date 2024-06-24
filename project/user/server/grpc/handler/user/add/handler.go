package add

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/proto/userpb"
	"hwCalendar/user/model/user"
	"hwCalendar/user/storage"
)

func Handle(ctx context.Context, req *userpb.AddUserRequest) (*userpb.AddUserResponse, error) {
	newUser, err := user.New(
		req.Username,
		req.Password,
	)
	if err != nil {
		return nil, handleError(err)
	}

	id, err := newUser.Add(ctx)
	if err != nil {
		return nil, handleError(err)
	}

	return &userpb.AddUserResponse{
		Id: int32(id),
	}, nil
}

func handleError(err error) error {
	if errors.Is(err, user.ErrEmptyUsername) ||
		errors.Is(err, user.ErrEmptyPassword) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, storage.ErrAlreadyExists) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
