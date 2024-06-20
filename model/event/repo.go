package event

import (
	"database/sql"
	"errors"
	"hwCalendar/storage"
)

func ById(id int) (*Event, error) {
	var event Event
	err := pgStorage.QueryRow("SELECT id, name, description, start_time FROM events WHERE id = $1", id).
		Scan(&event.Id, &event.Name, &event.Description, &event.Timestamp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, err
	}

	return &event, nil
}

func All() ([]Event, error) {
	events := make([]Event, 0)
	rows, err := pgStorage.Query("SELECT id, name, description, start_time FROM events")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Timestamp); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func Delete(id int) error {
	deletedEvent := &Event{Id: id}
	return deletedEvent.Delete()
}
