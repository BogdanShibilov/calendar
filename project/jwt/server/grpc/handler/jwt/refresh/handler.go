package refresh

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/jwt/model/jwt"
	"hwCalendar/jwt/server/grpc/handler/jwt/common"
	"hwCalendar/jwt/transport/grpc/user"
	"hwCalendar/proto/jwtpb"
	"hwCalendar/proto/userpb"
)

func Handle(ctx context.Context, req *jwtpb.RefreshTokensRequest) (*jwtpb.RefreshTokensResponse, error) {
	err := validate(req)
	if err != nil {
		return nil, err
	}

	claims, err := jwt.ParseRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, handleError(err)
	}

	if !jwt.IsRefreshTokenInRedis(ctx, claims.Id, req.RefreshToken) {
		return nil, status.Error(codes.Unauthenticated, jwt.ErrNoSuchTokenForUser.Error())
	}

	res, err := user.GetClient().UserById(ctx, &userpb.UserByIdRequest{Id: int32(claims.Id)})
	if err != nil {
		return nil, handleError(err)
	}

	tokenPair, err := jwt.GeneratePair(ctx, int(res.User.Id), res.User.Username)
	if err != nil {
		return nil, handleError(err)
	}

	return &jwtpb.RefreshTokensResponse{
		Pair: &jwtpb.TokenPair{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
	}, nil
}

func validate(req *jwtpb.RefreshTokensRequest) error {
	if req.RefreshToken == "" {
		return status.Error(codes.InvalidArgument, common.ErrMissingTokenString.Error())
	}

	return nil
}

func handleError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}
	if errors.Is(err, jwt.ErrMalformedToken) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, jwt.ErrExpiredToken) ||
		errors.Is(err, jwt.ErrUnknownClaims) {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
