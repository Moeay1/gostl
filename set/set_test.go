package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetNew(t *testing.T) {
	set := New[int](1, 2, 3)
	assert.Equal(t, 3, set.Len())
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
}

func TestSetAdd(t *testing.T) {
	set := New[int]()
	set.Add(1, 2, 3)
	assert.Equal(t, 3, set.Len())
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
}

func TestSetForEach(t *testing.T) {
	set := New[int](1, 2, 3)
	var sum int
	set.ForEach(func(val int) {
		sum += val
	})
	assert.Equal(t, 6, sum)
}

func TestSetClear(t *testing.T) {
	set := New[int](1, 2, 3)
	set.Clear()
	assert.Equal(t, 0, set.Len())
	assert.False(t, set.Contains(1))
	assert.False(t, set.Contains(2))
	assert.False(t, set.Contains(3))
}

func TestSetContains(t *testing.T) {
	set := New[int](1, 2, 3)
	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(4))
}

func TestSetIter(t *testing.T) {
	set := New[int](1, 2, 3)
	var sum int
	for val := range set.Iter() {
		sum += val
	}
	assert.Equal(t, 6, sum)
}

func TestSetLen(t *testing.T) {
	set := New[int](1, 2, 3)
	assert.Equal(t, 3, set.Len())
}

func TestSetDel(t *testing.T) {
	set := New[int](1, 2, 3)
	set.Del(1)
	assert.Equal(t, 2, set.Len())
	assert.False(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
}

func TestSetFuncUnion(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	set3 := Union[int](set1, set2)
	assert.Equal(t, 4, set3.Len())
	assert.True(t, set3.Contains(1))
	assert.True(t, set3.Contains(2))
	assert.True(t, set3.Contains(3))
	assert.True(t, set3.Contains(4))
}

func TestSetFuncDifference(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	set3 := Difference[int](set1, set2)
	assert.Equal(t, 1, set3.Len())
	assert.True(t, set3.Contains(1))
	assert.False(t, set3.Contains(2))
	assert.False(t, set3.Contains(3))
	assert.False(t, set3.Contains(4))
}

func TestSetFuncIntersect(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	set3 := Intersect[int](set1, set2)
	assert.Equal(t, 2, set3.Len())
	assert.False(t, set3.Contains(1))
	assert.True(t, set3.Contains(2))
	assert.True(t, set3.Contains(3))
	assert.False(t, set3.Contains(4))
}

func TestSetMethodUnion(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	set3 := set1.Union(set2)
	assert.Equal(t, 4, set3.Len())
	assert.True(t, set3.Contains(1))
	assert.True(t, set3.Contains(2))
	assert.True(t, set3.Contains(3))
	assert.True(t, set3.Contains(4))
}

func TestSetMethodDifference(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	set3 := set1.Difference(set2)
	assert.Equal(t, 1, set3.Len())
	assert.True(t, set3.Contains(1))
	assert.False(t, set3.Contains(2))
	assert.False(t, set3.Contains(3))
	assert.False(t, set3.Contains(4))
}

func TestSetMethodIntersect(t *testing.T) {
	set1 := New[int](1, 2, 3)
	set2 := New[int](2, 3, 4)
	set3 := set1.Intersect(set2)
	assert.Equal(t, 2, set3.Len())
	assert.False(t, set3.Contains(1))
	assert.True(t, set3.Contains(2))
	assert.True(t, set3.Contains(3))
	assert.False(t, set3.Contains(4))
}
