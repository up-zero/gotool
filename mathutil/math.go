package mathutil

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Abs 绝对值
func Abs[T Number](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

// Min 最小值
func Min[T Number](arg ...T) T {
	res := arg[0]
	for _, v := range arg {
		if v < res {
			res = v
		}
	}
	return res
}

// Max 最大值
func Max[T Number](arg ...T) T {
	res := arg[0]
	for _, v := range arg {
		if v > res {
			res = v
		}
	}
	return res
}
