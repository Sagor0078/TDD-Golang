package Sync

// concurrent counter with sync.Mutex and sync waitGroup

import (
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}


func NewSafeCounter() *SafeCounter {
	return &SafeCounter{v : make(map[string]int)}
}

func (c * SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

func (c * SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

