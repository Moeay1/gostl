package list

import (
	"github.com/Moeay1/gostl/i"
)

type List[T comparable] []T

func New[T comparable](v ...T) *List[T] {
	list := make(List[T], 0, len(v))
	list.Add(v...)
	return &list
}

func NewByLen[T comparable](length int, v ...T) *List[T] {
	list := make(List[T], 0, length)
	list.Add(v...)
	return &list
}

func NewByIter[T comparable](iter i.Iter[T]) *List[T] {
	list := New[T]()
	for val := range iter.Iter() {
		list.Add(val)
	}
	return list
}

func NewByLenIter[T comparable](lenIter i.LenIter[T]) *List[T] {
	list := NewByLen[T](lenIter.Len())
	for val := range lenIter.Iter() {
		list.Add(val)
	}
	return list
}

func (l *List[T]) Add(items ...T) {
	*l = append(*l, items...)
}

func (l *List[T]) Clear() {
	*l = make([]T, 0)
}

func (l *List[T]) Contains(val T) bool {
	for i := range *l {
		if val == (*l)[i] {
			return true
		}
	}
	return false
}

func (l *List[T]) ToSlice() []T {
	return []T(*l)
}

func (l *List[T]) Reverse() {
	for j, k := 0, len(*l)-1; j < k; j, k = j+1, k-1 {
		(*l)[j], (*l)[k] = (*l)[k], (*l)[j]
	}
}

func (l *List[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for _, e := range *l {
			ch <- e
		}
		close(ch)
	}()
	return ch
}

func (l *List[T]) Len() int {
	return len(*l)
}

func (l *List[T]) Insert(i int, val T) {
	newList := append((*l)[:i], append([]T{val}, (*l)[i:]...)...)
	*l = newList
}

func (l *List[T]) Del(i int) {
	newList := append((*l)[:i], (*l)[i+1:]...)
	*l = newList
}
