package main

import (
	"sync"
)

// Counter struct with a mutex and a value.
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter creates and returns a new instance of Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter's value.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
