package add

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/model/event"
	"hwCalendar/protobuf/eventpb"
	"hwCalendar/storage"
	"math/rand"
)

func Handle(_ context.Context, req *eventpb.AddEventRequest) (*eventpb.AddEventResponse, error) {
	id, err := generateUniqId()
	if err != nil {
		return nil, err
	}

	newEvent := eventFromReq(req, id)

	id, err = newEvent.Add()
	if err != nil {
		return nil, handleError(err)
	}

	return &eventpb.AddEventResponse{
		Id: int32(id),
	}, nil
}

func eventFromReq(req *eventpb.AddEventRequest, id int) *event.Event {
	return event.New(
		id,
		req.Name,
		req.Description,
		req.Timestamp.AsTime(),
	)
}

func generateUniqId() (int, error) {
	var randId int
	for {
		randId = rand.Intn(2147483647)
		_, err := event.ById(randId)
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				return randId, nil
			}

			return -1, status.Errorf(codes.Internal, "add event failed: %v", err)
		}
	}
}

func handleError(err error) error {
	if errors.Is(err, event.ErrInvalidId) ||
		errors.Is(err, event.ErrEmptyName) ||
		errors.Is(err, event.ErrEmptyDescription) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, storage.ErrAlreadyExists) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	if errors.Is(err, storage.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Errorf(codes.Internal, "add event failed: %v", err)
}
