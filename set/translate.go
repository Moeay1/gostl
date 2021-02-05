package set

import (
	"errors"

	"github.com/Moeay1/gostl"
)

func ToSet(iter gostl.IterAble) (*Set, error) {
	set := New()
	for e := range iter.Iter() {
		if t, ok := e.(gostl.HashAble); ok {
			set.Add(t)
		} else {
			return nil, errors.New("type not iterable")
		}
	}
	return set, nil
}
