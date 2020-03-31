package lru

import (
	"container/list"
)

// Cache cache data struct
type Cache struct {
	maxBytes  int64
	usedBytes int64
	ll        *list.List
	hashMap   map[string]*list.Element
	OnEvicted func(key string, value string)
}

type Container interface {
	Get(key string) string
	RemoveOldest()
	Add(key string, value string)
}

type Entry struct {
	key   string
	value string
}

// New function
func New(maxBytes int64, onEvicted func(key string, value string)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		hashMap:   make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value string, ok bool) {
	if e, ok := c.hashMap[key]; ok {
		c.ll.MoveToFront(e)
		value := e.Value.(*Entry).value
		return value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	e := c.ll.Back()
	if e != nil {
		c.ll.Remove(e)
		kv := e.Value.(*Entry)
		delete(c.hashMap, kv.key)
		c.usedBytes -= int64(len(kv.key) + len(kv.value))
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value string) {
	// update the existing item & move front
	if e, ok := c.hashMap[key]; ok {
		c.ll.MoveToFront(e)
		kv := e.Value.(*Entry)
		c.usedBytes += int64(len(value) - len(kv.value))
		kv.value = value
	} else {
		// create a brand-new entry
		e := c.ll.PushFront(&Entry{key, value})
		c.usedBytes += int64(len(key) + len(value))
		c.hashMap[key] = e
	}
	for c.maxBytes != 0 && c.usedBytes > c.maxBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
