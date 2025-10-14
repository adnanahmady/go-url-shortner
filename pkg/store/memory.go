package store

type StoreManager interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Has(key string) bool
}

type MemoryStoreManager struct {
	urls map[string]string
}

func NewMemoryStore() *MemoryStoreManager {
	return &MemoryStoreManager{urls: make(map[string]string)}
}

func (s *MemoryStoreManager) Has(key string) bool {
	_, ok := s.urls[key]
	return ok
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