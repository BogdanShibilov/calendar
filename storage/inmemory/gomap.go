package inmemory

import (
	"hwCalendar/storage"
	"sync"
)

const (
	EventType = "event"
)

var (
	once   sync.Once
	single *MapStorage
)

type MapStorage struct {
	m  map[string]map[any]any
	mu sync.RWMutex
}

func GetMapStorage() *MapStorage {
	once.Do(func() {
		single = &MapStorage{
			m: make(map[string]map[any]any),
		}
		single.m[EventType] = make(map[any]any)
	})

	return single
}

func (s *MapStorage) Add(modelType string, key, value any) (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.isKnownModelType(modelType) {
		return nil, ErrUnknownType
	}

	if isExistingKey(s.m[modelType], key) {
		return nil, storage.ErrAlreadyExists
	}

	s.m[modelType][key] = value

	return key, nil
}

func (s *MapStorage) ByKey(modelType string, key any) (any, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.isKnownModelType(modelType) {
		return nil, ErrUnknownType
	}

	value, ok := s.m[modelType][key]
	if !ok {
		return nil, storage.ErrNotFound
	}

	return value, nil
}

func (s *MapStorage) All(modelType string) ([]any, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.isKnownModelType(modelType) {
		return nil, ErrUnknownType
	}

	typeMap := s.m[modelType]
	valuesCount := len(typeMap)
	values := make([]any, 0, valuesCount)
	for _, v := range typeMap {
		values = append(values, v)
	}

	return values, nil
}

func (s *MapStorage) Update(modelType string, key, value any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isKnownModelType(modelType) {
		return ErrUnknownType
	}

	if !isExistingKey(s.m[modelType], key) {
		return storage.ErrNotFound
	}

	s.m[modelType][key] = value

	return nil
}

func (s *MapStorage) Delete(modelType string, key any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isKnownModelType(modelType) {
		return ErrUnknownType
	}

	delete(s.m[modelType], key)
	return nil
}

func (s *MapStorage) isKnownModelType(modelType string) bool {
	_, ok := s.m[modelType]
	return ok
}

func isExistingKey(m map[any]any, key any) bool {
	_, ok := m[key]
	return ok
}
