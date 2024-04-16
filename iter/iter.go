package iter

import (
	"github.com/Moeay1/gostl/i"
)

type Iter[T comparable] struct {
	ch <-chan T
}

func New[T comparable](iter i.Iter[T]) *Iter[T] {
	return &Iter[T]{ch: iter.Iter()}
}

func (i *Iter[T]) Iter() <-chan T {
	return i.ch
}

func (i *Iter[T]) Filter(filter func(T) bool) *Iter[T] {
	ch := make(chan T)
	go func() {
		for val := range i.ch {
			if filter(val) {
				ch <- val
			}
		}
		close(ch)
	}()
	return &Iter[T]{ch: ch}
}

func (i *Iter[T]) Map(mapper func(T) T) *Iter[T] {
	ch := make(chan T)
	go func() {
		for val := range i.ch {
			ch <- mapper(val)
		}
		close(ch)
	}()
	return &Iter[T]{ch: ch}
}

func (i *Iter[T]) Reduce(reducer func(T, T) T) T {
	var result T
	for val := range i.ch {
		result = reducer(result, val)
	}
	return result
}

func (i *Iter[T]) ForEach(f func(T)) {
	for val := range i.ch {
		f(val)
	}
}
