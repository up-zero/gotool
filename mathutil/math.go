package mathutil

import "github.com/up-zero/gotool"

// Abs 绝对值
func Abs[T gotool.Number](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

// Min 最小值
func Min[T gotool.Number](arg ...T) T {
	res := arg[0]
	for _, v := range arg {
		if v < res {
			res = v
		}
	}
	return res
}

// Max 最大值
func Max[T gotool.Number](arg ...T) T {
	res := arg[0]
	for _, v := range arg {
		if v > res {
			res = v
		}
	}
	return res
}

// Sum 求和
func Sum[T gotool.Number](arg ...T) T {
	var sum T
	for _, v := range arg {
		sum += v
	}
	return sum
}

// Average 平均值
func Average[T gotool.Number](arg ...T) T {
	return Sum(arg...) / T(len(arg))
}
