package event

import "hwCalendar/storage/inmemory"

func ById(id int) (*Event, error) {
	event, err := storage.ByKey(inmemory.EventType, id)
	if err != nil {
		return nil, err
	}

	return event.(*Event), nil
}

func All() ([]Event, error) {
	anySlice, err := storage.All(inmemory.EventType)
	if err != nil {
		return nil, err
	}

	events := make([]Event, len(anySlice))
	for i, e := range anySlice {
		events[i] = *e.(*Event)
	}
	return events, nil
}

func Delete(id int) error {
	deletedEvent := &Event{Id: id}
	return deletedEvent.Delete()
}
