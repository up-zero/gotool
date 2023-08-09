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
