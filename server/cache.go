package server

import (
	"time"
)

type CachingItem struct {
	CreatedAt time.Time
	TTL       time.Duration
	Data      interface{}
}

type CachingStore struct {
	DefaultTTL  time.Duration
	Data        *map[string]CachingItem
	LastRenewed time.Time
}

func NewCachingStore(defaultTTL time.Duration) *CachingStore {
	return &CachingStore{
		Data:       &map[string]CachingItem{},
		DefaultTTL: defaultTTL,
	}
}

// Get returns the value stored in the cache for a key, or nil if no value is present.
func (c *CachingStore) Get(key string) (interface{}, bool) {
	item, ok := (*c.Data)[key]
	if !ok {
		return nil, false
	}

	// check if item is expired
	if item.TTL != 0 && time.Since(item.CreatedAt) > item.TTL {
		// delete item
		c.Delete(key)
		return nil, false
	}

	return item.Data, true
}

// Set stores a value for a key with a TTL.
func (c *CachingStore) Set(key string, data interface{}, ttl time.Duration) {
	if ttl == 0 {
		ttl = c.DefaultTTL
	}

	(*c.Data)[key] = CachingItem{
		CreatedAt: time.Now(),
		TTL:       ttl,
		Data:      data,
	}
}

// Delete removes a key from the cache.
func (c *CachingStore) Delete(key string) {
	delete(*c.Data, key)
}
