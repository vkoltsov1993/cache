package cache

type Cache struct {
	cacheMap map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cacheMap: make(map[string]interface{}),
	}
}

func (cache *Cache) Set(key string, value interface{}) {
	cache.cacheMap[key] = value
}

func (cache Cache) Get(key string) interface{} {
	return cache.cacheMap[key]
}

func (cache *Cache) Delete(key string) bool {
	_, exists := cache.cacheMap[key]
	if exists {
		delete(cache.cacheMap, key)
		return true
	}
	return false
}
