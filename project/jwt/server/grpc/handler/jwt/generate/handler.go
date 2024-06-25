package generate

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/jwt/model/jwt"
	"hwCalendar/jwt/server/grpc/handler/jwt/common"
	"hwCalendar/proto/jwtpb"
)

func Handle(ctx context.Context, req *jwtpb.GenerateTokensRequest) (*jwtpb.GenerateTokensResponse, error) {
	err := validate(req)
	if err != nil {
		return nil, err
	}

	tokenPair, err := jwt.GeneratePair(ctx, int(req.Id), req.Username)
	if err != nil {
		return nil, handleError(err)
	}

	return &jwtpb.GenerateTokensResponse{
		Pair: &jwtpb.TokenPair{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
	}, nil
}

func validate(req *jwtpb.GenerateTokensRequest) error {
	if req.Username == "" {
		return status.Errorf(codes.InvalidArgument, common.ErrMissingUsername.Error())
	}

	return nil
}

func handleError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Errorf(codes.DeadlineExceeded, "deadline exceeded")
	}

	return status.Errorf(codes.Internal, err.Error())
}
