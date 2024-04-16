package i

type Pair[K comparable, V comparable] struct {
	First  K
	Second V
}

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
	ForEach(func(val T))
}

type Set[T comparable] interface {
	Container[T]

	Del(val ...T)
	Union(s Set[T]) Set[T]
	Difference(s Set[T]) Set[T]
	Intersect(s Set[T]) Set[T]
}

type List[T comparable] interface {
	Container[T]

	Insert(i int, val T)
	Del(i int)
	Reverse()
	ToSlice() []T
	Sort(func(a, b T) bool)
	Get(i int) T
	Set(i int, val T)
	AddAll(iter Iter[T])
}

type Map[K comparable, V comparable] interface {
	Lener

	ContainsKey(key K) bool
	Get(key K) V
	Set(key K, val V)
	Del(key K) V
	KeySet() Set[K]
	Iter[Pair[K, V]]
}
