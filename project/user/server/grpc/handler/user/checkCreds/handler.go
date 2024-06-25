package checkCreds

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hwCalendar/proto/userpb"
	"hwCalendar/user/model/user"
	"hwCalendar/user/server/grpc/handler/user/common"
	"hwCalendar/user/storage"
)

func Handle(ctx context.Context, req *userpb.CheckCredentialsRequest) (*userpb.CheckCredentialsResponse, error) {
	err := validate(req)
	if err != nil {
		return nil, err
	}

	u, err := user.ByUsername(ctx, req.Username)
	if err != nil {
		return nil, handleError(err)
	}

	err = u.VerifyPassword(req.Password)
	if err != nil {
		return nil, handleError(err)
	}

	return &userpb.CheckCredentialsResponse{
		User: &userpb.User{
			Id:        int32(u.Id),
			Username:  u.Username,
			PassHash:  u.PassHash,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
		},
	}, nil
}

func handleError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}
	if errors.Is(err, storage.ErrNotFound) ||
		errors.Is(err, user.ErrInvalidPassword) {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}

func validate(req *userpb.CheckCredentialsRequest) error {
	if req.Username == "" {
		return status.Error(codes.InvalidArgument, common.ErrMissingUsername.Error())
	}
	if req.Password == "" {
		return status.Error(codes.InvalidArgument, common.ErrMissingPassword.Error())
	}
	return nil
}
