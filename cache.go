package cache

import (
	"errors"
	"fmt"
)

type Cache struct {
	cacheMap map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cacheMap: make(map[string]interface{}),
	}
}

func (cache *Cache) Set(key string, value interface{}) error {
	_, exists := cache.cacheMap[key]
	if exists {
		errorMessage := fmt.Sprintf("%s key has been already added to cache", key)
		return errors.New(errorMessage)
	}
	cache.cacheMap[key] = value
	return nil
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
