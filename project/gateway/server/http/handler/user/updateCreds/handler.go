package updateCreds

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/user"
	"hwCalendar/proto/userpb"
	"net/http"
)

type Request struct {
	NewUsername string `json:"newUsername"`
	NewPassword string `json:"newPassword"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var req *Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handlerError(w, err)
		return
	}

	err = validate(req)
	if err != nil {
		handlerError(w, err)
	}

	err = updateUser(r.Context(), req, r.Context().Value("userId").(int32))
	if err != nil {
		handlerError(w, err)
		return
	}
}

func validate(req *Request) error {
	if req.NewUsername == "" {
		return ErrMissingNewUsername
	}
	if req.NewPassword == "" {
		return ErrMissingNewPassword
	}
	return nil
}

func updateUser(ctx context.Context, req *Request, userId int32) error {
	_, err := user.GetClient().UpdateUser(
		ctx,
		&userpb.UpdateUserRequest{
			Id:       userId,
			Username: req.NewUsername,
			Password: req.NewPassword,
		},
	)
	return err
}

func handlerError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.DeadlineExceeded:
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
		case codes.NotFound:
			http.Error(w, ErrAccessTokenForNonExistingUser.Error(), http.StatusNotFound)
		default:
			http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
		}
		return
	}
	if errors.Is(err, ErrMissingNewUsername) ||
		errors.Is(err, ErrMissingNewPassword) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}
