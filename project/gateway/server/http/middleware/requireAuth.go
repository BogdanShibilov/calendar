package middleware

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

var (
	ErrMissingAuthorizationHeader = errors.New("authorization header is required")
	ErrInvalidToken               = errors.New("authorization token is invalid")
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := getToken(r)
		if err != nil {
			handleError(w, err)
			return
		}

		parsedToken, err := parseAccessToken(r.Context(), tokenString)
		if err != nil {
			handleError(w, err)
		}

		r = r.WithContext(context.WithValue(r.Context(), "userId", parsedToken.UserId))
		r = r.WithContext(context.WithValue(r.Context(), "username", parsedToken.Username))

		next.ServeHTTP(w, r)
	}
}

func parseAccessToken(ctx context.Context, accessToken string) (*jwtpb.ParseAccessTokenResponse, error) {
	return jwt.GetClient().ParseAccessToken(
		ctx,
		&jwtpb.ParseAccessTokenRequest{
			AccessToken: accessToken,
		},
	)
}

func getToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrMissingAuthorizationHeader
	}
	tokens := strings.Split(authHeader, " ")
	if len(tokens) != 2 || tokens[0] != "Bearer" {
		return "", ErrInvalidToken
	}
	return tokens[1], nil
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

	if errors.Is(err, ErrMissingAuthorizationHeader) ||
		errors.Is(err, ErrInvalidToken) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}
