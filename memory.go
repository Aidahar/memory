package memory

type Memory struct {
	cache map[string]interface{}
}

func New() *Memory {
	return &Memory{
		cache: make(map[string]interface{}),
	}
}

func (m *Memory) Set(key string, value interface{}) {
	m.cache[key] = value
}

func (m *Memory) Get(key string) interface{} {
	return m.cache[key]
}
func (m *Memory) Delete(key string) {
	delete(m.cache, key)
}
