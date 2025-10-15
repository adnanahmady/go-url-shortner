package store

import (
	"sync"
)

type StoreManager interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Has(key string) bool
	HasUrl(url string) (string, bool)
	Lock()
	Unlock()
}

type MemoryStoreManager struct {
	mu *sync.Mutex
	urls map[string]string
}

func NewMemoryStore() *MemoryStoreManager {
	return &MemoryStoreManager{
		mu: &sync.Mutex{},
		urls: make(map[string]string),
	}
}

func (s *MemoryStoreManager) Has(key string) bool {
	_, ok := s.urls[key]
	return ok
}

func (s *MemoryStoreManager) HasUrl(url string) (string, bool) {
	for s, v := range s.urls {
		if v == url {
			return s, true
		}
	}
	return "", false
}

func (s *MemoryStoreManager) Set(key string, value string) error {
	if s.Has(key) {
		return ErrKeyAlreadyExists
	}
	s.urls[key] = value
	return nil
}

func (s *MemoryStoreManager) Get(key string) (string, error) {
	if !s.Has(key) {
		return "", ErrKeyNotFound
	}
	return s.urls[key], nil
}

func (s *MemoryStoreManager) Lock() {
	s.mu.Lock()
}

func (s *MemoryStoreManager) Unlock() {
	s.mu.Unlock()
}
