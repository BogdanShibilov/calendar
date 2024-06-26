package event

import (
	"context"
	"database/sql"
	"errors"
	"hwCalendar/calendar/storage"
	"hwCalendar/calendar/storage/postgres"
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
	var insertedId int
	err := pgStorage.QueryRowxContext(
		ctx,
		"INSERT INTO events (name, description, start_time) VALUES ($1, $2, $3) RETURNING id",
		e.Name, e.Description, e.Timestamp,
	).Scan(&insertedId)
	if err != nil {
		return 0, err
	}

	e.Id = insertedId

	return insertedId, nil
}

func (e *Event) Update(ctx context.Context, newName, newDesc string) error {
	e.Name = newName
	e.Description = newDesc

	_, err := pgStorage.NamedExecContext(
		ctx,
		"UPDATE events SET name = :name, description = :description WHERE id = :id",
		e,
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
	_, err := pgStorage.NamedExecContext(ctx, "DELETE FROM events WHERE id = :id", e)
	if err != nil {
		return err
	}

	return nil
}
