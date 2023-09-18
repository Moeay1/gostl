package set

import (
	"github.com/Moeay1/gostl/i"
)

func Union[T comparable](s1 i.Set[T], s2 i.Set[T]) i.Set[T] {
	length := s1.Len()
	if s2.Len() > length {
		length = s2.Len()
	}
	newSet := NewByLen[T](length)
	for e := range s1.Iter() {
		newSet.Add(e)
	}
	for e := range s2.Iter() {
		newSet.Add(e)
	}
	return newSet
}

func Difference[T comparable](s1 i.Set[T], s2 i.Set[T]) i.Set[T] {
	newSet := New[T]()
	for e := range s1.Iter() {
		if !s2.Contains(e) {
			newSet.Add(e)
		}
	}
	return newSet
}

func Intersect[T comparable](s1 i.Set[T], s2 i.Set[T]) i.Set[T] {
	newSet := New[T]()

	if s1.Len() < s2.Len() {
		for e := range s1.Iter() {
			if s2.Contains(e) {
				newSet.Add(e)
			}
		}
	} else {
		for e := range s2.Iter() {
			if s1.Contains(e) {
				newSet.Add(e)
			}
		}
	}
	return newSet
}
