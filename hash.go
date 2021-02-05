package gostl

// HashAble 可 hash 用于去重
type HashAble interface {
	Hash() string
}
