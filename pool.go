package just

import (
	"sync"
)

type Pool[T any] struct {
	pool  sync.Pool
	reset func(T)
}

// NewPool return a new pool with concrete type and reset fn.
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

// Get return another object from pool.
func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

// Put will reset object and put object to pool.
func (p *Pool[T]) Put(obj T) {
	p.reset(obj)
	p.pool.Put(obj)
}
