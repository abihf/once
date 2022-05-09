package once

import "sync"

type LoadFunc[T any] func() (T, error)

type Once[T any] struct {
	once sync.Once
	load LoadFunc[T]
	val  T
	err  error
}

func New[T any](load LoadFunc[T]) *Once[T] {
	return &Once[T]{
		load: load,
	}
}

func (o *Once[T]) Get() (T, error) {
	o.once.Do(func() {
		o.val, o.err = o.load()
	})
	return o.val, o.err
}
