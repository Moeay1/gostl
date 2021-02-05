package set

import (
	"sync"

	"github.com/Moeay1/gostl"
)

// Set 集合 需要实现 HashAble 接口 Set是线程安全的
type Set struct {
	lock  sync.RWMutex
	items map[string]gostl.HashAble
}

func New(items ...gostl.HashAble) *Set {
	set := Set{
		lock:  sync.RWMutex{},
		items: make(map[string]gostl.HashAble),
	}
	for i := 0; i < len(items); i++ {
		set.items[items[i].Hash()] = items[i]
	}
	return &set
}

// Add 添加元素，无论是否有该元素
func (s *Set) Add(items ...gostl.HashAble) {
	s.lock.Lock()
	for i := 0; i < len(items); i++ {
		s.items[items[i].Hash()] = items[i]
	}
	s.lock.Unlock()
}

// Discard 删除元素，无论是否有该元素
func (s *Set) Discard(items ...gostl.HashAble) {
	s.lock.Lock()
	for i := 0; i < len(items); i++ {
		delete(s.items, items[i].Hash())
	}
	s.lock.Unlock()
}

// Has 判断是否有某元素
func (s *Set) Exist(h gostl.HashAble) bool {
	s.lock.RLock()
	_, ok := s.items[h.Hash()]
	s.lock.RUnlock()
	return ok
}

// Iter 遍历
func (s *Set) Iter() <-chan interface{} {
	iter := make(chan interface{})
	go func() {
		s.lock.RLock()
		for _, e := range s.items {
			iter <- e
		}
		s.lock.RUnlock()
		close(iter)
	}()
	return iter
}

// Len 长度
func (s *Set) Len() int {
	s.lock.RLock()
	length := len(s.items)
	s.lock.RUnlock()
	return length
}
