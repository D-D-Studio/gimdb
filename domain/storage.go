package domain

import (
	"errors"
	"sync"
)

type Storage struct {
	memory *sync.Map
}

func NewStorage() *Storage {
	storage := new(Storage)

	storage.memory = new(sync.Map)

	return storage
}

func (s *Storage) Set(key string, value string) {
	s.memory.Store(key, value)
}

func (s *Storage) Get(key string) (string, error) {
	value, ok := s.memory.Load(key)

	if !ok {
		return "", errors.New("cannot find value by key")
	}

	return value.(string), nil
}

func (s *Storage) Delete(key string) error {
	_, loaded := s.memory.LoadAndDelete(key)

	if !loaded {
		return errors.New("cannot find value by key")
	}

	return nil
}
