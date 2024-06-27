package refresh

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/jwt"
	"hwCalendar/proto/jwtpb"
	"net/http"
)

type Request struct {
	RefreshToken string `json:"refresh_token"`
}

type Response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var req *Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, err)
		return
	}

	err = validate(req)
	if err != nil {
		handleError(w, err)
		return
	}

	refreshTokensRes, err := refreshTokens(r.Context(), req)
	if err != nil {
		handleError(w, err)
		return
	}

	res := &Response{
		AccessToken:  refreshTokensRes.Pair.AccessToken,
		RefreshToken: refreshTokensRes.Pair.RefreshToken,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		handleError(w, err)
		return
	}
}

func refreshTokens(ctx context.Context, req *Request) (*jwtpb.RefreshTokensResponse, error) {
	return jwt.GetClient().RefreshTokens(
		ctx,
		&jwtpb.RefreshTokensRequest{
			RefreshToken: req.RefreshToken,
		},
	)
}

func validate(req *Request) error {
	if req.RefreshToken == "" {
		return ErrMissingRefreshToken
	}
	return nil
}

func handleError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.DeadlineExceeded:
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
		case codes.InvalidArgument:
			http.Error(w, ErrInvalidToken.Error(), http.StatusBadRequest)
		case codes.Unauthenticated:
			http.Error(w, ErrUnauthenticated.Error(), http.StatusUnauthorized)
		default:
			http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
		}
		return
	}

	if errors.Is(err, ErrMissingRefreshToken) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}
