package inmemory

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hwCalendar/model"
	"hwCalendar/storage"
	"sync"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Run("should add new event", func(t *testing.T) {
		s := events{
			m: make(map[int]model.Event),
		}

		assert.Equal(t, 0, len(s.m))
		simpleEvent := model.Event{Id: 111, Name: "Simple", Description: "Very simple", Timestamp: time.Now()}
		simpleId, err := s.Add(simpleEvent)
		require.NoError(t, err)
		assert.Equal(t, 111, simpleId)
		assert.Equal(t, 1, len(s.m))
	})

	t.Run("should return error that event already exists", func(t *testing.T) {
		simpleEvent := model.Event{Id: 111, Name: "Simple", Description: "Very simple", Timestamp: time.Now()}
		s := events{
			m: map[int]model.Event{111: simpleEvent},
		}

		_, err := s.Add(simpleEvent)
		assert.ErrorIs(t, err, storage.ErrAlreadyExists)
	})
}

func TestById(t *testing.T) {
	t.Run("should get event by id", func(t *testing.T) {
		simpleEvent := model.Event{Id: 111, Name: "Simple", Description: "Very simple", Timestamp: time.Now()}
		s := events{
			m: map[int]model.Event{111: simpleEvent},
		}

		e, err := s.ById(111)
		require.NoError(t, err)
		assert.Equal(t, simpleEvent, *e)
	})

	t.Run("should return error that event not found", func(t *testing.T) {
		s := events{
			m: make(map[int]model.Event),
		}

		_, err := s.ById(111)
		assert.ErrorIs(t, err, storage.ErrNotFound)
	})
}

func TestAll(t *testing.T) {
	t.Run("should get all events", func(t *testing.T) {
		simpleEvent := model.Event{Id: 111, Name: "Simple", Description: "Very simple", Timestamp: time.Now()}
		awesomeEvent := model.Event{Id: 222, Name: "Awesome", Description: "Very awesome", Timestamp: time.Now()}
		s := events{
			m: map[int]model.Event{111: simpleEvent, 222: awesomeEvent},
		}

		e := s.All()
		assert.Equal(t, 2, len(e))
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update event", func(t *testing.T) {
		simpleEvent := model.Event{Id: 111, Name: "Simple", Description: "Very simple", Timestamp: time.Now()}
		awesomeEvent := model.Event{Id: 222, Name: "Awesome", Description: "Very awesome", Timestamp: time.Now()}
		s := events{
			m: map[int]model.Event{111: simpleEvent, 222: awesomeEvent},
		}

		newName := "Yet another name"
		simpleEvent.Name = newName
		err := s.Update(simpleEvent)
		require.NoError(t, err)
		e, _ := s.ById(111)
		assert.Equal(t, newName, e.Name)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should delete event", func(t *testing.T) {
		simpleEvent := model.Event{Id: 111, Name: "Simple", Description: "Very simple", Timestamp: time.Now()}
		awesomeEvent := model.Event{Id: 222, Name: "Awesome", Description: "Very awesome", Timestamp: time.Now()}
		s := events{
			m: map[int]model.Event{111: simpleEvent, 222: awesomeEvent},
		}

		s.Delete(222)
		assert.Equal(t, 1, len(s.m))
	})
}

func TestConcurrency(t *testing.T) {
	t.Run("should add many new events while using goroutines", func(t *testing.T) {
		s := events{
			m: make(map[int]model.Event),
		}

		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, _ = s.Add(model.Event{
					Id:          i,
					Name:        "Name",
					Description: "Description",
					Timestamp:   time.Now(),
				})
			}()
		}
		wg.Wait()
		assert.Equal(t, 1000, len(s.m))
	})
}

func TestSingleton(t *testing.T) {
	t.Run("constructor should always return the same storage", func(t *testing.T) {
		firstCall := NewEvents()
		secondCall := NewEvents()
		assert.Same(t, firstCall, secondCall)
	})
}
