package all

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/gateway/server/http/handler/common"
	"hwCalendar/gateway/transport/grpc/calendar"
	"hwCalendar/proto/eventpb"
	"net/http"
	"time"
)

type EventDTO struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
}

type Response struct {
	Events []EventDTO
}

func Handle(w http.ResponseWriter, r *http.Request) {
	allEventsRes, err := getAllEvents(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	res := resFrom(allEventsRes)
	err = json.NewEncoder(w).Encode(res)
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

func getAllEvents(ctx context.Context) (*eventpb.AllEventsResponse, error) {
	return calendar.GetClient().AllEvents(
		ctx,
		&emptypb.Empty{},
	)
}

func resFrom(pbRes *eventpb.AllEventsResponse) *Response {
	result := make([]EventDTO, 0, len(pbRes.Events))
	for _, event := range pbRes.Events {
		result = append(result, EventDTO{
			Id:          int(event.Id),
			Name:        event.Name,
			Description: event.Description,
			StartTime:   event.Timestamp.AsTime(),
		})
	}

	return &Response{
		Events: result,
	}
}
