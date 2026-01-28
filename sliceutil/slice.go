package sliceutil

import (
	"fmt"
	"strconv"
	"strings"

	"slices"

	"github.com/up-zero/gotool"
)

// Unique 切片去重，求并集
//
// # Params:
//
//	elems: 传入同种类型切片
func Unique[T comparable](elems ...[]T) []T {
	var ans = make([]T, 0)
	var m = make(map[T]struct{})
	for _, elem := range elems {
		for _, v := range elem {
			if _, ok := m[v]; !ok {
				ans = append(ans, v)
			}
			m[v] = struct{}{}
		}
	}
	return ans
}

// Contains 切片是否包含某个值
//
// # Params:
//
//	s: 待遍历的切片
//	target: 目标值
func Contains[T gotool.Number | string | bool](s []T, target T) bool {
	return slices.Contains(s, target)
}

// Join 切片拼接成字符串
//
// # Params:
//
//	elems: 待拼接的数值
//	sep: 拼接用的字符串
func Join[T gotool.Number | string](elems []T, sep string) string {
	if len(elems) == 0 {
		return ""
	}
	var ans strings.Builder
	for i, elem := range elems {
		if i > 0 {
			ans.WriteString(sep)
		}
		switch v := any(elem).(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			ans.WriteString(fmt.Sprint(v))
		case float32:
			ans.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 32))
		case float64:
			ans.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		case string:
			ans.WriteString(v)
		default:
			ans.WriteString(fmt.Sprint(v))
		}
	}
	return ans.String()
}

// Concat 切片拼接
func Concat[T any](elems ...[]T) []T {
	var l int
	for _, elem := range elems {
		l += len(elem)
	}
	var ans = make([]T, 0, l)
	for _, elem := range elems {
		ans = append(ans, elem...)
	}
	return ans
}

// Intersect 求多个切片的交集，元素存在时会先去重
//
//	elems: 多个切片
//
// # Example:
//
//	Intersect([]int{12, 22, 12}, []int{12, 22})  // [12 22]
func Intersect[T comparable](elems ...[]T) []T {
	if len(elems) == 0 {
		return []T{}
	}

	var common = make(map[T]struct{})
	for _, v := range elems[0] {
		common[v] = struct{}{}
	}
	for i := 1; i < len(elems); i++ {
		set := make(map[T]struct{})
		for _, v := range elems[i] {
			set[v] = struct{}{}
		}
		for k := range common {
			if _, ok := set[k]; !ok {
				delete(common, k)
			}
		}
		if len(common) == 0 {
			return []T{}
		}
	}

	var ans = make([]T, 0)
	for _, v := range elems[0] {
		if _, ok := common[v]; ok {
			ans = append(ans, v)
			delete(common, v)
		}
	}
	return ans
}

// Filter 切片过滤
//
// # Params:
//
//	s: 待过滤的切片
//	f: 过滤函数
func Filter[T any](s []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Map 切片类型转换
//
// # Params:
//
//	s: 待转换的切片
//	iteratee: 转换函数
func Map[T any, R any](s []T, iteratee func(item T) R) []R {
	result := make([]R, len(s))
	for i, item := range s {
		result[i] = iteratee(item)
	}
	return result
}

// GroupBy 将切片按指定的 Key 进行分组
//
// # Params:
//
//	s: 待分组的切片
//	iteratee: 分组函数
func GroupBy[T any, K comparable](s []T, iteratee func(item T) K) map[K][]T {
	res := make(map[K][]T)
	for _, v := range s {
		key := iteratee(v)
		res[key] = append(res[key], v)
	}
	return res
}

// Chunk 将切片按指定大小切分为多个小切片
//
// # Params:
//
//	s: 待切分的切片
//	size: 切片大小
func Chunk[T any](s []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	var res [][]T
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		res = append(res, s[i:end])
	}
	return res
}

// Difference 求差集 (s1 - s2)，返回在 s1 中但不在 s2 中的元素
func Difference[T comparable](s1, s2 []T) []T {
	m := make(map[T]struct{}, len(s2))
	for _, v := range s2 {
		m[v] = struct{}{}
	}
	res := make([]T, 0)
	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
