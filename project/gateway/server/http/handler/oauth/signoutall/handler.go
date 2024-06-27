package signoutall

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/jwt"
	"hwCalendar/proto/jwtpb"
	"net/http"
	"strings"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(
		r.Header.Get("Authorization"),
		" ",
	)[1]

	err := signOutAll(r.Context(), token)
	if err != nil {
		handleError(w, err)
		return
	}
}

func signOutAll(ctx context.Context, accessToken string) error {
	_, err := jwt.GetClient().RemoveAllTokensForUser(
		ctx,
		&jwtpb.RemoveAllTokensForUserRequest{
			AccessToken: accessToken,
		},
	)
	return err
}

func handleError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.DeadlineExceeded:
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
		case codes.Unauthenticated:
			http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
		default:
			http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
		}
		return
	}

	if errors.Is(err, ErrMissingAccessToken) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}
