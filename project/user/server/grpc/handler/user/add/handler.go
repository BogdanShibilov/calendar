package add

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/proto/userpb"
	"hwCalendar/user/model/user"
	"hwCalendar/user/server/grpc/handler/user/common"
	"hwCalendar/user/storage"
)

func Handle(ctx context.Context, req *userpb.AddUserRequest) (*userpb.AddUserResponse, error) {
	err := validate(req)
	if err != nil {
		return nil, err
	}

	_, err = user.ByUsername(ctx, req.Username)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

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

func validate(req *userpb.AddUserRequest) error {
	if req.Username == "" {
		return status.Error(codes.InvalidArgument, common.ErrMissingUsername.Error())
	}
	if req.Password == "" {
		return status.Error(codes.InvalidArgument, common.ErrMissingPassword.Error())
	}
	return nil
}

func handleError(err error) error {
	if errors.Is(err, storage.ErrAlreadyExists) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
