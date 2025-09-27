package postgres

import (
	"container/list"
	"errors"
	"order_service/internal/model"
	"sync"
)

var defaultCacheSizeError = errors.New("defaultCacheSize должен быть > 0")
var startCacheSizeError = errors.New("startCacheSize должен быть >= 0")
var startMustBeGteDefaultSizeError = errors.New("startCacheSize должен быть >= defaultCacheSize")

type cacheEntry struct {
	key   string
	value *model.Order
}

type LRUCache struct {
	mu        sync.Mutex
	maxSize   int
	evictList *list.List               // Для вытеснения
	cache     map[string]*list.Element // Для доступа
}

// NewLRUCache создает новый LRU-кэш с заданным размером
func NewLRUCache(size int) (*LRUCache, error) {
	if size <= 0 {
		return nil, defaultCacheSizeError
	}
	return &LRUCache{
		maxSize:   size,
		evictList: list.New(),
		cache:     make(map[string]*list.Element),
	}, nil
}

// Put добавляет элемент в кэш
func (c *LRUCache) Put(key string, value *model.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Если элемент уже есть в кэше
	if element, ok := c.cache[key]; ok {
		c.evictList.MoveToFront(element)
		element.Value.(*cacheEntry).value = value
		return
	}

	// Если элемент новый, добавляем его в начало списка
	entry := &cacheEntry{key, value}
	element := c.evictList.PushFront(entry)
	c.cache[key] = element

	// Если кэш переполнен, удаляем самый старый элемент
	if c.evictList.Len() > c.maxSize {
		c.removeOldest()
	}
}

// Get получает значение из кэша
func (c *LRUCache) Get(key string) (*model.Order, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.cache[key]; ok {
		c.evictList.MoveToFront(element)
		return element.Value.(*cacheEntry).value, true
	}

	return nil, false
}

// removeOldest - метод для удаления самого старого элемента
func (c *LRUCache) removeOldest() {
	element := c.evictList.Back()
	if element != nil {
		c.evictList.Remove(element)
		entry := element.Value.(*cacheEntry)
		delete(c.cache, entry.key)
	}
}
