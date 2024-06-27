package byid

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/calendar"
	"hwCalendar/proto/eventpb"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	idstring := r.URL.Query().Get("id")
	if idstring == "" {
		http.Error(w, "id query parameter is required", http.StatusBadRequest)
	}
	id, err := strconv.Atoi(idstring)
	if err != nil {
		http.Error(w, "id query parameter must be an integer", http.StatusBadRequest)
	}
	eventRes, err := getEventById(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}

	res := &Response{
		Id:          int(eventRes.Event.Id),
		Name:        eventRes.Event.Name,
		Description: eventRes.Event.Description,
		StartTime:   eventRes.Event.Timestamp.AsTime(),
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
		case codes.NotFound:
			http.Error(w, ErrEventNotFound.Error(), http.StatusNotFound)
		case codes.DeadlineExceeded:
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
		default:
			http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}

func getEventById(ctx context.Context, id int) (*eventpb.EventByIdResponse, error) {
	return calendar.GetClient().EventById(
		ctx,
		&eventpb.EventByIdRequest{
			Id: int32(id),
		},
	)
}
