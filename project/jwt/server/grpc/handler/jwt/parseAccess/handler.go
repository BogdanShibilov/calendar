package parseAccess

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/jwt/model/jwt"
	"hwCalendar/proto/jwtpb"
)

func Handle(ctx context.Context, req *jwtpb.ParseAccessTokenRequest) (*jwtpb.ParseAccessTokenResponse, error) {
	claims, err := jwt.ParseAccessToken(req.AccessToken)
	if err != nil {
		return nil, handleError(err)
	}

	if !jwt.IsAccessTokenInRedis(ctx, claims.UserId, claims.ID, req.AccessToken) {
		return nil, status.Error(codes.Unauthenticated, jwt.ErrNoSuchTokenForUser.Error())
	}

	return &jwtpb.ParseAccessTokenResponse{
		UserId:   int32(claims.UserId),
		Username: claims.Username,
	}, nil
}

func handleError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}
	if errors.Is(err, jwt.ErrExpiredToken) ||
		errors.Is(err, jwt.ErrUnknownClaims) ||
		errors.Is(err, jwt.ErrMalformedToken) {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
