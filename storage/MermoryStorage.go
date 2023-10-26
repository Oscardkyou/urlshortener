package storage

import (
	"errors"
	"sync"
)

type MemoryStorage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}

func (m *MemoryStorage) Save(shortURL, longURL string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.data[shortURL]; exists {
		return errors.New("shortURL already exists")
	}

	m.data[shortURL] = longURL
	return nil
}

func (m *MemoryStorage) Resolve(shortURL string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	longURL, exists := m.data[shortURL]
	if !exists {
		return "", errors.New("shortURL not found")
	}

	return longURL, nil
}
