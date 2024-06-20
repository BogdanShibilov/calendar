package event

import (
	"database/sql"
	"errors"
	"hwCalendar/storage"
	"hwCalendar/storage/postgres"
	"time"
)

var pgStorage = postgres.GetDb()

type Event struct {
	Id          int
	Name        string
	Description string
	Timestamp   time.Time
}

func New(name string, desc string, timestamp time.Time) *Event {
	return &Event{
		Name:        name,
		Description: desc,
		Timestamp:   timestamp,
	}
}

func (e *Event) Add() (int, error) {
	if err := validate(e); err != nil {
		return -1, err
	}

	var insertedId int
	err := pgStorage.QueryRow(
		"INSERT INTO events (name, description, start_time) VALUES ($1, $2, $3) RETURNING id",
		e.Name, e.Description, e.Timestamp,
	).Scan(&insertedId)
	if err != nil {
		return -1, err
	}

	return insertedId, nil
}

func (e *Event) Update(newName, newDesc string) error {
	e.Name = newName
	e.Description = newDesc

	if err := validate(e); err != nil {
		return err
	}

	_, err := pgStorage.Exec(
		"UPDATE events SET name = $1, description = $2 WHERE id = $3",
		e.Name, e.Description, e.Id,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrNotFound
		}
		return err
	}

	return nil
}

func (e *Event) Delete() error {
	_, err := pgStorage.Exec("DELETE FROM events WHERE id = $1", e.Id)
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
