package signup

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
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	CreatedUserId int `json:"created_user_id"`
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

	addUserRes, err := addUser(r.Context(), req)
	if err != nil {
		handleError(w, err)
		return
	}

	res := &Response{
		CreatedUserId: int(addUserRes.Id),
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		handleError(w, err)
		return
	}
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

func handleError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.AlreadyExists:
			http.Error(w, ErrUsernameAlreadyExists.Error(), http.StatusConflict)
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

func addUser(ctx context.Context, req *Request) (*userpb.AddUserResponse, error) {
	return user.GetClient().AddUser(
		ctx,
		&userpb.AddUserRequest{
			Username: req.Username,
			Password: req.Password,
		},
	)
}
