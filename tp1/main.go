package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cache := NewCache()
	cache.Set("users:123", "Alice", 2*time.Second)
	if val, ok := cache.Get("users:123"); ok {
		fmt.Println(val)
	} else {
		fmt.Print("Pas trouvé ou expiré")
	}
	time.Sleep(2 * time.Second)
	if val, ok := cache.Get("users:123"); ok {
		fmt.Println(val)
	} else {
		fmt.Print("Pas trouvé ou expiré 2 eme tentative")
	}

}

type CacheItem struct {
	Value      interface{}
	Expiration int64
}
type Cache struct {
	data map[string]CacheItem
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]CacheItem)}
}
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
}
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.data[key]
	if !found || time.Now().UnixNano() > item.Expiration {
		return nil, false
	}
	return item.Value, true
}
