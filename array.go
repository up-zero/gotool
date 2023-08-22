package gotool

// ArrayDuplicate 数组去重
//
// elems 传入同种类型数组
func ArrayDuplicate[T int64 | int32 | int16 | int8 | int | string](elems ...[]T) []T {
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
func ArrayIn[T int64 | int32 | int16 | int8 | int | uint | uint8 | uint16 | uint32 | uint64 | string |
	bool | float64 | float32](target T, arr []T) bool {
	for _, v := range arr {
		if target == v {
			return true
		}
	}
	return false
}
