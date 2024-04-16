package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNew(t *testing.T) {
	list := New[int](1, 2, 3)
	assert.Equal(t, []int{1, 2, 3}, list.ToSlice())
}

func TestListNewByLength(t *testing.T) {
	list := NewByLen[int](3)
	list.Add(1, 2, 3)
	assert.Equal(t, []int{1, 2, 3}, list.ToSlice())
}

func TestListAdd(t *testing.T) {
	list := New[int]()
	list.Add(1, 2, 3)
	assert.Equal(t, []int{1, 2, 3}, list.ToSlice())
}

func TestListForEach(t *testing.T) {
	list := New[int](1, 2, 3)
	var sum int
	list.ForEach(func(val int) {
		sum += val
	})
	assert.Equal(t, 6, sum)
}

func TestListAddAll(t *testing.T) {
	list := New[int]()
	list.AddAll(New[int](1, 2, 3))
	assert.Equal(t, []int{1, 2, 3}, list.ToSlice())
}

func TestListIter(t *testing.T) {
	list := New[int](1, 2, 3)
	var sum int
	for val := range list.Iter() {
		sum += val
	}
	assert.Equal(t, 6, sum)
}

func TestListLen(t *testing.T) {
	list := New[int](1, 2, 3)
	assert.Equal(t, 3, list.Len())
}

func TestListClear(t *testing.T) {
	list := New[int](1, 2, 3)
	list.Clear()
	assert.Equal(t, 0, list.Len())
}

func TestListContains(t *testing.T) {
	list := New[int](1, 2, 3)
	assert.True(t, list.Contains(1))
	assert.False(t, list.Contains(4))
}

func TestListDel(t *testing.T) {
	list := New[int](1, 2, 3)
	list.Del(1)
	assert.Equal(t, []int{1, 3}, list.ToSlice())
}

func TestListInsert(t *testing.T) {
	list := New[int](1, 2, 3)
	list.Insert(1, 4)
	assert.Equal(t, []int{1, 4, 2, 3}, list.ToSlice())
}

func TestListReverse(t *testing.T) {
	list := New[int](1, 2, 3)
	list.Reverse()
	assert.Equal(t, []int{3, 2, 1}, list.ToSlice())
}

func TestListSort(t *testing.T) {
	list := New[int](3, 2, 1)
	list.Sort(func(a, b int) bool {
		return a < b
	})
	assert.Equal(t, []int{1, 2, 3}, list.ToSlice())
}

func TestListGet(t *testing.T) {
	list := New[int](1, 2, 3)
	assert.Equal(t, 2, list.Get(1))
}

func TestListSet(t *testing.T) {
	list := New[int](1, 2, 3)
	list.Set(1, 4)
	assert.Equal(t, []int{1, 4, 3}, list.ToSlice())
}

func TestListToMap(t *testing.T) {
	list := New[int](1, 2, 3)
	m := ToMap(list, func(val int) (int, int) {
		return val, val
	})
	assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 3}, m)
}

type User struct {
	ID   int
	Name string
}

func TestListToMapWithStruct(t *testing.T) {
	list := New[User](User{ID: 1, Name: "Alice"}, User{ID: 2, Name: "Bob"})
	m := ToMap(list, func(val User) (int, User) {
		return val.ID, val
	})
	assert.Equal(t, map[int]User{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}, m)
}

func TestIterFilter(t *testing.T) {
	data := New[int](1, 2, 3, 4, 5)
	f := func(v int) bool {
		return v%2 == 0
	}
	expected := New[int](2, 4)
	result := NewByIter[int](data.Stream().Filter(f))
	assert.Equal(t, expected, result)
}

func TestIterMap(t *testing.T) {
	data := New[int](1, 2, 3)
	m := func(v int) int {
		return v * 2
	}
	expected := New[int](2, 4, 6)
	result := NewByIter[int](data.Stream().Map(m))
	assert.Equal(t, expected, result)
}

func TestIterReduce(t *testing.T) {
	data := New[int](1, 2, 3)
	r := func(a, b int) int {
		return a + b
	}
	expected := 6
	result := data.Stream().Reduce(r)
	assert.Equal(t, expected, result)
}

func TestIterFilterMap(t *testing.T) {
	data := New[int](1, 2, 3, 4, 5)
	f := func(v int) bool {
		return v%2 == 0
	}
	m := func(v int) int {
		return v * 2
	}
	expected := New[int](4, 8)
	result := NewByIter[int](data.Stream().Filter(f).Map(m))
	assert.Equal(t, expected, result)
}

func TestIterMapFilter(t *testing.T) {
	data := New[int](1, 2, 3, 4, 5)
	m := func(v int) int {
		return v * 2
	}
	f := func(v int) bool {
		return v%4 == 0
	}
	expected := New[int](4, 8)
	result := NewByIter[int](data.Stream().Map(m).Filter(f))
	assert.Equal(t, expected, result)
}

func TestIterFilterReduce(t *testing.T) {
	data := New[int](1, 2, 3, 4, 5)
	f := func(v int) bool {
		return v%2 == 0
	}
	r := func(a, b int) int {
		return a + b
	}
	expected := 6
	result := data.Stream().Filter(f).Reduce(r)
	assert.Equal(t, expected, result)
}

func TestIterMapReduce(t *testing.T) {
	data := New[int](1, 2, 3)
	m := func(v int) int {
		return v * 2
	}
	r := func(a, b int) int {
		return a + b
	}
	expected := 12
	result := data.Stream().Map(m).Reduce(r)
	assert.Equal(t, expected, result)
}
