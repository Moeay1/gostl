package gostl

// LenAble 可计算长度
type LenAble interface {
	Len() int
}

func Len(l LenAble) int {
	return l.Len()
}
