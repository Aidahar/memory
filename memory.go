package cache

import "errors"

type Cache struct {
	cache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (m *Cache) Set(key string, value interface{}) {
	m.cache[key] = value
}

func (m *Cache) Get(key string) (interface{}, error) {
	k, ok := m.cache[key]
	if !ok {
		return nil, errors.New("invalid key")
	}
	return k, nil
}
func (m *Cache) Delete(key string) {
	delete(m.cache, key)
}
