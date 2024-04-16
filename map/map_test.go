package mmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	m := New[int, int]()
	assert.Equal(t, 0, m.Len())
}

func TestMapSet(t *testing.T) {
	m := New[int, int]()
	m.Set(1, 2)
	assert.Equal(t, 2, m.Get(1))
}

func TestMapDel(t *testing.T) {
	m := New[int, int]()
	m.Set(1, 2)
	assert.Equal(t, 2, m.Del(1))
}

func TestMapContainsKey(t *testing.T) {
	m := New[int, int]()
	m.Set(1, 2)
	assert.True(t, m.ContainsKey(1))
	assert.False(t, m.ContainsKey(2))
}

func TestMapKeySet(t *testing.T) {
	m := New[int, int]()
	m.Set(1, 2)
	m.Set(2, 3)
	keySet := m.KeySet()
	assert.True(t, keySet.Contains(1))
	assert.True(t, keySet.Contains(2))
	assert.False(t, keySet.Contains(3))
}

func TestMapIter(t *testing.T) {
	m := New[int, int]()
	m.Set(1, 2)
	m.Set(2, 3)
	sum := 0
	for pair := range m.Iter() {
		sum += pair.Second
	}
	assert.Equal(t, 5, sum)
}

func TestMapLen(t *testing.T) {
	m := New[int, int]()
	m.Set(1, 2)
	m.Set(2, 3)
	assert.Equal(t, 2, m.Len())
}
