package signin

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/jwt"
	"hwCalendar/gateway/transport/grpc/user"
	"hwCalendar/proto/jwtpb"
	"hwCalendar/proto/userpb"
	"net/http"
)

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	checkCredsRes, err := checkCreds(r.Context(), req)
	if err != nil {
		handleError(w, err)
		return
	}

	generateTokensRes, err := generateTokens(r.Context(), checkCredsRes)
	if err != nil {
		handleError(w, err)
		return
	}

	res := &Response{
		AccessToken:  generateTokensRes.Pair.AccessToken,
		RefreshToken: generateTokensRes.Pair.RefreshToken,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		handleError(w, err)
		return
	}
}

func handleError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.Unauthenticated:
			http.Error(w, ErrWrongCredentials.Error(), http.StatusUnauthorized)
		case codes.DeadlineExceeded:
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
		default:
			http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
		}
		return
	}

	if errors.Is(err, ErrMissingUsername) ||
		errors.Is(err, ErrMissingPassword) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}

func validate(req *Request) error {
	if req.Username == "" {
		return ErrMissingUsername
	}
	if req.Password == "" {
		return ErrMissingPassword
	}

	return nil
}

func checkCreds(ctx context.Context, req *Request) (*userpb.CheckCredentialsResponse, error) {
	return user.GetClient().CheckCredentials(
		ctx,
		&userpb.CheckCredentialsRequest{
			Username: req.Username,
			Password: req.Password,
		},
	)
}

func generateTokens(ctx context.Context, credsRes *userpb.CheckCredentialsResponse) (*jwtpb.GenerateTokensResponse, error) {
	return jwt.GetClient().GenerateTokens(
		ctx,
		&jwtpb.GenerateTokensRequest{
			UserId:   credsRes.User.Id,
			Username: credsRes.User.Username,
		},
	)
}
