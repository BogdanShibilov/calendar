package event

import (
	"hwCalendar/storage/inmemory"
	"time"
)

var storage = inmemory.GetMapStorage()

type Event struct {
	Id          int
	Name        string
	Description string
	Timestamp   time.Time
}

func (e *Event) Add() (int, error) {
	if err := validate(e); err != nil {
		return -1, err
	}

	id, err := storage.Add(inmemory.EventType, e.Id, e)
	if err != nil {
		return -1, err
	}

	return id.(int), nil
}

func (e *Event) Update() error {
	if err := validate(e); err != nil {
		return err
	}

	err := storage.Update(inmemory.EventType, e.Id, e)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete() error {
	err := storage.Delete(inmemory.EventType, e.Id)
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
