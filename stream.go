package gotool

type Stream[T any] struct {
	data []T
}

// NewStream 初始化 Stream
func NewStream[T any](data []T) *Stream[T] {
	return &Stream[T]{
		data: data,
	}
}

// Filter 数据过滤
//
// # Examples:
//
// s := NewStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//
// s.Filter(func(v int) bool { return v%2 == 0 }) // []int{2, 4, 6, 8, 10}
func (s *Stream[T]) Filter(f func(T) bool) *Stream[T] {
	result := make([]T, 0)
	for _, v := range s.data {
		if f(v) {
			result = append(result, v)
		}
	}
	s.data = result
	return s
}

// Map 数据处理
//
// # Examples:
//
// s := NewStream([]int{1, 2, 3, 4, 5})
//
// s.Map(func(v int) int { return v * 2 })
func (s *Stream[T]) Map(fn func(item T) T) *Stream[T] {
	return StreamMap[T](s, fn)
}

// StreamMap 数据处理与转换
//
// # Examples:
//
// s := NewStream([]int{1, 2, 3, 4, 5})
//
// StreamMap(s, func(v int) string { return fmt.Sprintf("data: %d", v) })
func StreamMap[T, R any](s *Stream[T], f func(T) R) *Stream[R] {
	result := make([]R, len(s.data))
	for i, v := range s.data {
		result[i] = f(v)
	}
	return &Stream[R]{
		data: result,
	}
}

// ToSlice 转换为切片
func (s *Stream[T]) ToSlice() []T {
	return s.data
}
