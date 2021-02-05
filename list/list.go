package list

import (
	"container/list"
	"sync"
)

type List struct {
	lock  sync.RWMutex
	items []interface{}
}

func New(items ...interface{}) *List {
	list := List{
		lock:  sync.RWMutex{},
		items: make([]interface{}, 0, len(items)),
	}
	list.items = append(list.items, items...)
	return &list
}

// Append 添加元素
func (l *List) Append(items ...interface{}) {
	l.lock.Lock()
	l.items = append(l.items, items...)
	l.lock.Unlock()
}

// Len 长度
func (l *List) Len() int {
	l.lock.RLock()
	length := len(l.items)
	l.lock.RUnlock()
	return length
}

// Iter 遍历
func (l *List) Iter() <-chan interface{} {
	iter := make(chan interface{})
	go func() {
		l.lock.RLock()
		for _, e := range l.items {
			iter <- e
		}
		l.lock.RUnlock()
		close(iter)
	}()
	return iter
}

// Swap 交换列表中两项
func (l *List) Swap(n, m int) {
	l.lock.Lock()
	l.items[n], l.items[m] = l.items[m], l.items[n]
	l.lock.Unlock()
}
