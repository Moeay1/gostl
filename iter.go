package gostl

// IterAble 迭代器，可以通过 range 来遍历
type IterAble interface {
	Iter() chan<- interface{}
}
