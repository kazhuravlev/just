package just

import (
	"sync"
)

type Pool[T any] struct {
	pool  sync.Pool
	reset func(T)
}

// NewPool returns a new pool with concrete type and resets fn.
func NewPool[T any](constructor func() T, reset func(T)) *Pool[T] {
	if reset == nil {
		reset = func(T) {}
	}

	return &Pool[T]{
		pool: sync.Pool{
			New: func() any { return constructor() },
		},
		reset: reset,
	}
}

// Get return another object from the pool.
func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

// Put will reset the object and put this object to the pool.
func (p *Pool[T]) Put(obj T) {
	p.reset(obj)
	p.pool.Put(obj)
}
