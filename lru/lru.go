package lru

import "container/list"

type Entry struct {
	Key   any
	Value any
}

type Cache struct {
	maxLength int

	list *list.List

	cache     map[any]*list.Element
	OnEvicted func(key any, value any)
}

func NewCache(maxLength int) *Cache {
	return &Cache{
		maxLength: maxLength,
		list:      list.New(),
		cache:     make(map[any]*list.Element),
	}
}

func (c *Cache) Add(key any, val any) {
	if element, ok := c.cache[key]; ok {
		c.list.MoveToFront(element)
		element.Value.(*Entry).Value = key
		return
	}

	frontElement := c.list.PushFront(&Entry{
		Key: key, Value: val,
	})
	c.cache[key] = frontElement

	if c.maxLength != 0 && c.list.Len() > c.maxLength {
		c.removeOldest()
	}
}

func (c *Cache) Get(key any) (any, bool) {
	if c.cache == nil {
		return nil, false
	}
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		entry := ele.Value.(*Entry)
		return entry.Value, ok
	}
	return nil, false
}

func (c *Cache) Remove(key any) {

	if c.cache == nil {
		return
	}
	if ele, ok := c.cache[key]; ok {
		c.removeElement(ele)
	}
}

func (c *Cache) Clear() {
	if c.OnEvicted != nil {
		front := c.list.Front()
		for front != nil {
			entry := front.Value.(*Entry)
			c.OnEvicted(entry.Key, entry.Value)
		}
	}
	c.list = nil
	c.cache = nil
}

func (c *Cache) removeElement(element *list.Element) {
	c.list.Remove(element)
	var ele = element.Value.(*Entry)
	delete(c.cache, ele.Key)
	if c.OnEvicted != nil {
		c.OnEvicted(ele.Key, ele.Value)
	}
}

func (c *Cache) removeOldest() {
	if c.cache == nil {
		return
	}
	back := c.list.Back()
	if back != nil {
		c.removeElement(back)
	}
}
