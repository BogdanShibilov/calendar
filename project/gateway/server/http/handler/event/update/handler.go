package update

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/calendar"
	"hwCalendar/proto/eventpb"
	"net/http"
)

type Request struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var req *Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, err)
		return
	}

	err = updateEvent(r.Context(), req)
	if err != nil {
		handleError(w, err)
		return
	}
}

func handleError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.DeadlineExceeded:
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
		case codes.NotFound:
			http.Error(w, ErrEventNotFound.Error(), http.StatusNotFound)
		default:
			http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}

func updateEvent(ctx context.Context, req *Request) error {
	_, err := calendar.GetClient().UpdateEvent(
		ctx,
		&eventpb.UpdateEventRequest{
			Id:          int32(req.Id),
			Name:        req.Name,
			Description: req.Description,
		},
	)
	return err
}
