package all

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

type EventDTO struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
}

type Response struct {
	Events     []EventDTO
	TotalPages int `json:"total_pages"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	limitString := r.URL.Query().Get("limit")
	pageString := r.URL.Query().Get("page")
	if limitString == "" {
		limitString = "10"
	}
	if pageString == "" {
		pageString = "1"
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		http.Error(w, "invalid limit query parameter", http.StatusBadRequest)
	}
	page, err := strconv.Atoi(pageString)
	if err != nil {
		http.Error(w, "invalid page query parameter", http.StatusBadRequest)
	}

	allEventsRes, err := getAllEvents(r.Context(), limit, page)
	if err != nil {
		handleError(w, err)
		return
	}

	res := Response{
		Events:     dtoFrom(allEventsRes),
		TotalPages: int(allEventsRes.TotalPages),
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
		case codes.DeadlineExceeded:
			http.Error(w, common.ErrRequestTimeout.Error(), http.StatusRequestTimeout)
		case codes.InvalidArgument:
			http.Error(w, "page is out of bounds", http.StatusBadRequest)
		}
		return
	}
	http.Error(w, common.ErrInternal.Error(), http.StatusInternalServerError)
}

func getAllEvents(ctx context.Context, limit, page int) (*eventpb.AllEventsResponse, error) {
	return calendar.GetClient().AllEvents(
		ctx,
		&eventpb.AllEventsRequest{
			Limit: int32(limit),
			Page:  int32(page),
		},
	)
}

func dtoFrom(pbRes *eventpb.AllEventsResponse) []EventDTO {
	result := make([]EventDTO, 0, len(pbRes.Events))
	for _, event := range pbRes.Events {
		result = append(result, EventDTO{
			Id:          int(event.Id),
			Name:        event.Name,
			Description: event.Description,
			StartTime:   event.Timestamp.AsTime(),
		})
	}

	return result
}
