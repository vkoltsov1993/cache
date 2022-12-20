package cache

type Cache struct {
	key   string
	value interface{}
}

func New() *Cache {
	return &Cache{
		key:   "",
		value: nil,
	}
}

type CacheBag struct {
}

func (cache *Cache) Set(key string, value interface{}) {
	ca
}
