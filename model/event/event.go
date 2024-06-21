package event

import (
	"context"
	"database/sql"
	"errors"
	"hwCalendar/storage"
	"hwCalendar/storage/postgres"
	"time"
)

var pgStorage = postgres.GetDb()

type Event struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Timestamp   time.Time `db:"start_time"`
}

func New(name string, desc string, timestamp time.Time) *Event {
	return &Event{
		Name:        name,
		Description: desc,
		Timestamp:   timestamp,
	}
}

func (e *Event) Add(ctx context.Context) (int, error) {
	if err := validate(e); err != nil {
		return -1, err
	}

	var insertedId int
	err := pgStorage.QueryRowxContext(
		ctx,
		"INSERT INTO events (name, description, start_time) VALUES ($1, $2, $3) RETURNING id",
		e.Name, e.Description, e.Timestamp,
	).Scan(&insertedId)
	if err != nil {
		return -1, err
	}

	e.Id = insertedId

	return insertedId, nil
}

func (e *Event) Update(ctx context.Context, newName, newDesc string) error {
	e.Name = newName
	e.Description = newDesc

	if err := validate(e); err != nil {
		return err
	}

	_, err := pgStorage.ExecContext(
		ctx,
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

func (e *Event) Delete(ctx context.Context) error {
	_, err := pgStorage.ExecContext(ctx, "DELETE FROM events WHERE id = $1", e.Id)
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
