package removeAllTokens

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/jwt/model/jwt"
	"hwCalendar/jwt/server/grpc/handler/jwt/common"
	"hwCalendar/proto/jwtpb"
)

func Handle(ctx context.Context, req *jwtpb.RemoveAllTokensForUserRequest) (*emptypb.Empty, error) {
	err := validate(req)
	if err != nil {
		return nil, err
	}

	claims, err := jwt.ParseAccessToken(req.AccessToken)
	if err != nil {
		return nil, handleError(err)
	}

	jwt.RemoveAllTokensFor(ctx, claims.UserId)

	return &emptypb.Empty{}, nil
}

func handleError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}
	if errors.Is(err, jwt.ErrExpiredToken) ||
		errors.Is(err, jwt.ErrMalformedToken) ||
		errors.Is(err, jwt.ErrUnknownClaims) {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}

func validate(req *jwtpb.RemoveAllTokensForUserRequest) error {
	if req.AccessToken == "" {
		return status.Error(codes.InvalidArgument, common.ErrMissingTokenString.Error())
	}

	return nil
}
