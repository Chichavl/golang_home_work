package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	_, found := l.items[key]
	// exist in cache
	if found {
		l.items[key].Value = cacheItem{key: key, value: value}
		l.queue.MoveToFront(l.items[key])
	}
	// new item
	if !found {
		l.items[key] = l.queue.PushFront(cacheItem{key: key, value: value})
	}
	// Aging cache
	if l.capacity < l.queue.Len() {
		ci, ok := l.queue.Back().Value.(cacheItem)
		if !ok {
			panic("Casting to type cacheItem failed")
		}

		delete(l.items, ci.key)
		l.queue.Remove(l.queue.Back())
	}

	return found
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	valueContainer, found := l.items[key]

	var val interface{}

	if valueContainer != nil {
		ci, ok := valueContainer.Value.(cacheItem)
		if !ok {
			panic("Casting to type cacheItem failed")
		}

		val = ci.value
	} else {
		val = valueContainer
	}

	if found {
		l.queue.MoveToFront(l.items[key])
	}

	return val, found
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache { //nolint: ireturn
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
