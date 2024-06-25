package update

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/calendar/storage"
	"hwCalendar/proto/userpb"
	"hwCalendar/user/model/user"
	"hwCalendar/user/storage/postgres"
)

var pgStorage = postgres.GetDb()

func Handle(ctx context.Context, req *userpb.UpdateUserRequest) (*emptypb.Empty, error) {
	id := int(req.Id)
	u, err := user.ById(ctx, id)
	if err != nil {
		return nil, handleError(err)
	}

	tx, err := pgStorage.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, handleError(err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	err = u.UpdateUsername(ctx, tx, req.Username)
	if err != nil {
		return nil, handleError(err)
	}
	err = u.UpdatePassword(ctx, tx, req.Password)
	if err != nil {
		return nil, handleError(err)
	}

	_ = tx.Commit()
	return &emptypb.Empty{}, nil
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
