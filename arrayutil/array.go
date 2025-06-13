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
// elems 传入同种类型数组
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
// arr 待遍历的数组
// target 目标值
func Contains[T gotool.Number | string | bool](arr []T, target T) bool {
	return slices.Contains(arr, target)
}

// Join 数组拼接成字符串
//
// elems 待拼接的数值
// sep 拼接用的字符串
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
