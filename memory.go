package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]interface{}
	mu    *sync.Mutex
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
		mu:    new(sync.Mutex),
	}
}

func (m *Cache) Set(key string, value interface{}, ttl time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache[key] = value
	f := func() {
		delete(m.cache, key)
	}
	time.AfterFunc(ttl, f)
}

func (m *Cache) Get(key string) (interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	k, ok := m.cache[key]
	if !ok {
		return nil, errors.New("invalid key")
	}
	return k, nil
}

func (m *Cache) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.cache, key)
}
