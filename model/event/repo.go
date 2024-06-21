package event

import (
	"context"
	"database/sql"
	"errors"
	"hwCalendar/storage"
)

func ById(ctx context.Context, id int) (*Event, error) {
	var event Event
	err := pgStorage.QueryRowxContext(
		ctx,
		"SELECT id, name, description, start_time FROM events WHERE id = $1", id,
	).StructScan(&event)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, err
	}

	return &event, nil
}

func All(ctx context.Context) ([]Event, error) {
	events := make([]Event, 0)
	rows, err := pgStorage.QueryxContext(ctx, "SELECT id, name, description, start_time FROM events")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var event Event
		if err := rows.StructScan(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func Delete(ctx context.Context, id int) error {
	deletedEvent := &Event{Id: id}
	return deletedEvent.Delete(ctx)
}
