package inmemory

import (
	"hwCalendar/model"
	"hwCalendar/storage"
	"sync"
)

type events struct {
	m  map[int]model.Event
	mu sync.RWMutex
}

var single *events
var once sync.Once

func NewEvents() *events {
	once.Do(func() {
		single = &events{
			m: make(map[int]model.Event),
		}
	})

	return single
}

func (e *events) Add(event model.Event) (int, error) {
	e.mu.Lock()
	_, ok := e.m[event.Id]
	if ok {
		return -1, storage.ErrAlreadyExists
	}

	e.m[event.Id] = event
	e.mu.Unlock()

	return event.Id, nil
}

func (e *events) ById(id int) (*model.Event, error) {
	e.mu.RLock()
	event, ok := e.m[id]
	e.mu.RUnlock()
	if !ok {
		return nil, storage.ErrNotFound
	}

	return &event, nil
}

func (e *events) All() []model.Event {
	eventCount := len(e.m)
	events := make([]model.Event, 0, eventCount)

	e.mu.RLock()
	for _, event := range e.m {
		events = append(events, event)
	}
	e.mu.RUnlock()

	return events
}

func (e *events) Update(newEvent model.Event) error {
	e.mu.Lock()
	_, ok := e.m[newEvent.Id]
	if !ok {
		return storage.ErrNotFound
	}

	e.m[newEvent.Id] = newEvent
	e.mu.Unlock()

	return nil
}

func (e *events) Delete(id int) {
	e.mu.Lock()
	delete(e.m, id)
	e.mu.Unlock()
}
