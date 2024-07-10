package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// Use channels when passing ownership of data
// Use mutexes for managing state

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
