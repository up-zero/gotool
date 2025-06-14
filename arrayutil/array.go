package arrayutil

import (
	"fmt"
	"strconv"
	"strings"

	"slices"

	"github.com/up-zero/gotool"
)

// Duplicate 数组去重
//
//	elems: 传入同种类型数组
func Duplicate[T gotool.Number | string](elems ...[]T) []T {
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

// Contains 数组是否包含某个值
//
//	arr: 待遍历的数组
//	target: 目标值
func Contains[T gotool.Number | string | bool](arr []T, target T) bool {
	return slices.Contains(arr, target)
}

// Join 数组拼接成字符串
//
//	elems: 待拼接的数值
//	sep: 拼接用的字符串
func Join[T gotool.Number | string](elems []T, sep string) string {
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

// Concat 数组合并
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

// Intersect 求多个数组的交集，元素存在时会先去重
//
//	elems: 多个数组
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
