package cache                                                                                                         
 
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

func (m *Cache) Get(key string) interface{} {
    k, exists := m.cache[key]
    if exists {
        return k
    }
    return nil

}
func (m *Cache) Delete(key string) {
    delete(m.cache, key)
}
