package deleteEvent

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/calendar"
	"hwCalendar/proto/eventpb"
	"net/http"
	"strconv"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	if idString == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "id query parameter must be a number", http.StatusBadRequest)
	}

	err = deleteEvent(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
}

func handleError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		if s.Code() == codes.DeadlineExceeded {
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
			return
		}
	}

	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}

func deleteEvent(ctx context.Context, id int) error {
	_, err := calendar.GetClient().DeleteEvent(
		ctx,
		&eventpb.DeleteEventRequest{
			Id: int32(id),
		},
	)
	return err
}
