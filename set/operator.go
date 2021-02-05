package set

import (
	"github.com/Moeay1/gostl"
)

func Union(a, b *Set) *Set {
	r := New()
	r.lock.Lock()
	a.lock.RLock()
	for k := range a.items {
		r.items[k] = a.items[k]
	}
	a.lock.RUnlock()
	b.lock.RLock()
	for k := range b.items {
		r.items[k] = b.items[k]
	}
	b.lock.RUnlock()
	r.lock.Unlock()
	return r
}

func Intersection(a, b *Set) *Set {
	r := New()
	r.lock.Lock()
	a.lock.RLock()
	b.lock.RLock()
	if gostl.Len(a) < gostl.Len(b) {
		for k := range a.items {
			if _, ok := b.items[k]; ok {
				r.items[k] = b.items[k]
			}
		}
	} else {
		for k := range b.items {
			if _, ok := a.items[k]; ok {
				r.items[k] = b.items[k]
			}
		}
	}
	a.lock.RUnlock()
	b.lock.RUnlock()
	r.lock.Unlock()
	return r
}
