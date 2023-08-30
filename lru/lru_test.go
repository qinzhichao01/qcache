package lru

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	cache := NewCache(4)
	cache.Add(1, 1)
	cache.Add(2, 2)
	cache.Add(3, 3)
	cache.Add(4, 4)
	cache.Add(5, 5)
	val, b := cache.Get(1)
	if b {
		fmt.Println(val)
	}
	get, b2 := cache.Get(2)
	if b2 {
		fmt.Println(get)

	}

}
