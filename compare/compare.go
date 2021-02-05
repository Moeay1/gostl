package compare

// Less 比较两个基本类型的大小，包括：
//  int8  int  int16  int32  int64
//  uint8 uint uint16 uint32 uint64
func Less(a, b interface{}) bool {
	switch m := a.(type) {
	case uint8:
		n, ok := b.(uint8)
		if !ok {
			panic("a type uint8 but b not type uint8")

		}
		return m < n
	case uint:
		n, ok := b.(uint)
		if !ok {
			panic("a type uint but b not type uint")
		}
		return m < n
	case uint16:
		n, ok := b.(uint16)
		if !ok {
			panic("a type uint16 but b not type uint16")
		}
		return m < n
	case uint32:
		n, ok := b.(uint32)
		if !ok {
			panic("a type uint32 but b not type uint32")
		}
		return m < n
	case uint64:
		n, ok := b.(uint64)
		if !ok {
			panic("a type uint64 but b not type uint64")
		}
		return m < n
	case int8:
		n, ok := b.(int8)
		if !ok {
			panic("a type int8 but b not type int8")
		}
		return m < n
	case int:
		n, ok := b.(int)
		if !ok {
			panic("a type int but b not type int")
		}
		return m < n
	case int16:
		n, ok := b.(int16)
		if !ok {
			panic("a type int16 but b not type int16")
		}
		return m < n
	case int32:
		n, ok := b.(int32)
		if !ok {
			panic("a type int32 but b not type int32")
		}
		return m < n
	case int64:
		n, ok := b.(int64)
		if !ok {
			panic("a type int64 but b not type int64")
		}
		return m < n
	default:
		panic("not suppport this type")
	}
}
