package cache

import (
	"errors"
	"time"
)

type Cache struct {
	cache     map[string]interface{}
	createdAt time.Time
	duration  time.Duration
}

func New() *Cache {
	return &Cache{
		cache:     make(map[string]interface{}),
		createdAt: time.Now(),
		duration:  0,
	}
}

func (m *Cache) Set(key string, value interface{}, ttl time.Duration) {
	m.cache[key] = value
	m.createdAt = time.Now()
	m.duration = ttl
}

func (m *Cache) Get(key string) (interface{}, error) {
	t := time.Now()
	k, ok := m.cache[key]
	if !ok {
		return nil, errors.New("invalid key")
	}
	if m.createdAt.Sub(t) < m.duration {
		delete(m.cache, key)
	}
	return k, nil
}

func (m *Cache) Delete(key string) {
	delete(m.cache, key)
}
