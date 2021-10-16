package memory

type Memory struct {
	Key   string
	Value interface{}
}

func (m *Memory) Set(key string, value interface{}) {
	m.Key = key
	m.Value = value
}

func (m *Memory) Get(key string) interface{} {
	return m.Value
}
func (m *Memory) Delete(key string) {
	m.Key = ""
}
