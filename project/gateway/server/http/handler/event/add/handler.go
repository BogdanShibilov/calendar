package add

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/calendar"
	"hwCalendar/proto/eventpb"
	"net/http"
	"time"
)

type Request struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
}

type Response struct {
	EventId int `json:"event_id"`
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
	}

	addEventRes, err := addEvent(r.Context(), req)
	if err != nil {
		handleError(w, err)
	}

	res := &Response{
		EventId: int(addEventRes.Id),
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		handleError(w, err)
		return
	}
}

func addEvent(ctx context.Context, req *Request) (*eventpb.AddEventResponse, error) {
	return calendar.GetClient().AddEvent(
		ctx,
		&eventpb.AddEventRequest{
			Name:        req.Name,
			Description: req.Description,
			Timestamp:   timestamppb.New(req.StartTime),
		},
	)
}

func validate(req *Request) error {
	if req.Name == "" {
		return ErrMissingName
	}
	if req.Description == "" {
		return ErrMissingDescription
	}

	return nil
}

func handleError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.AlreadyExists:
			http.Error(w, ErrEventAlreadyExists.Error(), http.StatusConflict)
		case codes.DeadlineExceeded:
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
		case codes.Internal:
			http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
		}
		return
	}

	if errors.Is(err, ErrMissingName) ||
		errors.Is(err, ErrMissingDescription) {
		http.Error(w, ErrMissingName.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}
