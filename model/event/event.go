package event

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwCalendar/storage"
	"hwCalendar/storage/inmemory"
	"math/rand"
	"time"
)

var mapStorage = inmemory.GetMapStorage()

type Event struct {
	Id          int
	Name        string
	Description string
	Timestamp   time.Time
}

func New(name string, desc string, timestamp time.Time) *Event {
	id, _ := generateUniqId()
	return &Event{
		Id:          id,
		Name:        name,
		Description: desc,
		Timestamp:   timestamp,
	}
}

func generateUniqId() (int, error) {
	var randId int
	for {
		randId = rand.Intn(2147483647)
		_, err := ById(randId)
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				return randId, nil
			}

			return -1, status.Errorf(codes.Internal, "add event failed: %v", err)
		}
	}
}

func (e *Event) Add() (int, error) {
	if err := validate(e); err != nil {
		return -1, err
	}

	id, err := mapStorage.Add(inmemory.EventType, e.Id, e)
	if err != nil {
		return -1, err
	}

	return id.(int), nil
}

func (e *Event) Update(newName, newDesc string) error {
	e.Name = newName
	e.Description = newDesc

	if err := validate(e); err != nil {
		return err
	}

	err := mapStorage.Update(inmemory.EventType, e.Id, e)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete() error {
	err := mapStorage.Delete(inmemory.EventType, e.Id)
	if err != nil {
		return err
	}

	return nil
}

func validate(event *Event) error {
	if event.Name == "" {
		return ErrEmptyName
	}
	if event.Description == "" {
		return ErrEmptyDescription
	}
	if event.Id < 0 {
		return ErrInvalidId
	}
	return nil
}
