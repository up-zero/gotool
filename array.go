package gotool

import (
	"reflect"
	"strconv"
)

// ArrayDuplicate 数组去重
//
// elems 传入同种类型数组
func ArrayDuplicate[T Number | string](elems ...[]T) []T {
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

// ArrayIn 数组是否包含某个值
//
// target 目标值
// arr 待遍历的数组
func ArrayIn[T Number | string | bool](target T, arr []T) bool {
	for _, v := range arr {
		if target == v {
			return true
		}
	}
	return false
}

// ArrayJoin 整型拼接
//
// elems 待拼接的数值
// sep 拼接用的字符串
func ArrayJoin[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](elems []T, sep string) string {
	var ans string
	for _, elem := range elems {
		switch reflect.TypeOf(elem).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			ans = ans + strconv.Itoa(int(elem)) + sep
		}
	}
	if ans == "" {
		return ""
	}
	return ans[:len(ans)-len(sep)]
}
