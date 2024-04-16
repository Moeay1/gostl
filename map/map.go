package mmap

import (
	"github.com/Moeay1/gostl/i"
	"github.com/Moeay1/gostl/set"
)

var _ i.Map[int, int] = (*Map[int, int])(nil)

type Map[K comparable, V comparable] map[K]V

func New[K comparable, V comparable]() *Map[K, V] {
	return &Map[K, V]{}
}

func (m *Map[K, V]) Len() int {
	return len(*m)
}

func (m *Map[K, V]) ContainsKey(key K) bool {
	_, isContains := (*m)[key]
	return isContains
}

func (m *Map[K, V]) Get(key K) V {
	return (*m)[key]
}

func (m *Map[K, V]) Set(key K, val V) {
	(*m)[key] = val
}

func (m *Map[K, V]) Del(key K) V {
	val := (*m)[key]
	delete(*m, key)
	return val
}

func (m *Map[K, V]) KeySet() i.Set[K] {
	set := set.New[K]()
	for key := range *m {
		set.Add(key)
	}
	return set
}

func (m *Map[K, V]) Iter() <-chan i.Pair[K, V] {
	ch := make(chan i.Pair[K, V])
	go func() {
		for key, val := range *m {
			ch <- i.Pair[K, V]{First: key, Second: val}
		}
		close(ch)
	}()
	return ch
}

func (m *Map[K, V]) KeyIter() <-chan K {
	ch := make(chan K)
	go func() {
		for key := range *m {
			ch <- key
		}
		close(ch)
	}()
	return ch
}

func (m *Map[K, V]) ValIter() <-chan V {
	ch := make(chan V)
	go func() {
		for _, val := range *m {
			ch <- val
		}
		close(ch)
	}()
	return ch
}
