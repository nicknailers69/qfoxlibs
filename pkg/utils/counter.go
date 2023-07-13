package utils

import "sync"

type Counter struct {
	mu  sync.Mutex
	cnt uint64
}

// Add increments the counter.
func (c *Counter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt++
}

// Val retrieves the counter's value.
func (c *Counter) Val() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	cnt := c.cnt
	return cnt
}

// counter is a thread-safe connection counter.
var SafeCounter Counter
