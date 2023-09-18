package set

import (
	"github.com/Moeay1/gostl/i"
)

type Set[T comparable] map[T]struct{}

func New[T comparable](v ...T) *Set[T] {
	set := make(Set[T], len(v))
	for _, val := range v {
		set.Add(val)
	}
	return &set
}

func NewByLen[T comparable](length int, v ...T) *Set[T] {
	set := make(Set[T], length)
	for _, val := range v {
		set.Add(val)
	}
	return &set
}

func NewByIter[T comparable](iter i.Iter[T]) *Set[T] {
	set := New[T]()
	for val := range iter.Iter() {
		set.Add(val)
	}
	return set
}

func (s *Set[T]) Add(v ...T) {
	for _, val := range v {
		(*s)[val] = struct{}{}
	}
}

func (s *Set[T]) Clear() {
	*s = make(Set[T])
}

func (s *Set[T]) Contains(val T) bool {
	_, isContains := (*s)[val]
	return isContains
}

func (s *Set[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for e := range *s {
			ch <- e
		}
		close(ch)
	}()
	return ch
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) Del(v ...T) {
	for _, val := range v {
		delete(*s, val)
	}
}
