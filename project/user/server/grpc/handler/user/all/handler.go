package all

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hwCalendar/proto/userpb"
	"hwCalendar/user/model/user"
)

func Handle(ctx context.Context, _ *emptypb.Empty) (*userpb.AllUsersResponse, error) {
	all, err := user.All(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Errorf(codes.DeadlineExceeded, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return resFromUsers(all), nil
}

func resFromUsers(allUsers []user.User) *userpb.AllUsersResponse {
	userPbSlice := make([]*userpb.User, 0, len(allUsers))
	for _, u := range allUsers {
		userPbSlice = append(userPbSlice, &userpb.User{
			Id:        int32(u.Id),
			Username:  u.Username,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
		})
	}

	return &userpb.AllUsersResponse{
		Users: userPbSlice,
	}
}
