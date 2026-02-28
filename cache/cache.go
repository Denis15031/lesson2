package main

import (
	"encoding/json"
	"sync"
	"time"
)

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]CacheItem
}

func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]CacheItem),
	}
	//фоновая очистка просроченных ключей каждую минуту
	go c.startCleanup()
	return c
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(), // время истечения
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key] //пытаемся найти элемент по ключу

	//если не найден или время истечения наступило > считаем просроченным
	if !found || time.Now().UnixNano() > item.Expiration {
		return nil, false
	}
	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *Cache) Exists(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	return found && time.Now().UnixNano() <= item.Expiration
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]CacheItem)
}

func (c *Cache) ToJSON() ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	valid := make(map[string]interface{})
	now := time.Now().UnixNano()

	for k, v := range c.items {
		if now <= v.Expiration {
			valid[k] = v.Value
		}
	}
	return json.Marshal(valid)
}

func GetAs[T any](c *Cache, key string) (T, error) {
	var zero T

	val, ok := c.Get(key)
	if !ok {
		return zero, &CacheError{"key not found or expired"}
	}
	if result, ok := val.(T); ok {
		return result, nil
	}
	return zero, &CacheError{"type assertion failed"}
}

type CacheError struct {
	msg string
}

func (e *CacheError) Error() string {
	return e.msg
}

func (c *Cache) startCleanup() {
	ticker := time.NewTicker(time.Minute)

	go func() {
		for range ticker.C {
			c.mu.Lock()
			now := time.Now().UnixNano()

			for k, v := range c.items {
				if now > v.Expiration {
					delete(c.items, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}
