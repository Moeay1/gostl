package i

// IterAble 迭代器，可以通过 range 来遍历
type Iter[T comparable] interface {
	Iter() <-chan T
}

// LenAble 可计算长度
type Lener interface {
	Len() int
}

type LenIter[T comparable] interface {
	Iter[T]
	Lener
}

func Len(l Lener) int {
	return l.Len()
}

type Container[T comparable] interface {
	Iter[T]
	Lener

	Add(val ...T)
	Clear()
	Contains(val T) bool
}

type Set[T comparable] interface {
	Container[T]

	Del(val ...T)
}

type List[T comparable] interface {
	Container[T]

	Insert(i int, val T)
	Del(i int)
	Reverse()
	ToSlice() []T
}
