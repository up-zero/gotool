package streamutil

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

// Extreme 返回流中的极值，比较函数决定是最大值还是最小值
//
// 要获取最大值，传入 func(a, b T) bool { return a > b }
//
// 要获取最小值，传入 func(a, b T) bool { return a < b }
//
// # Examples:
//
// s := NewStream([]int{1, 2, 3, 4, 5})
//
// max := s.Extreme(func(a, b int) bool { return a > b }) // 5
//
// min := s.Extreme(func(a, b int) bool { return a < b }) // 1
func (s *Stream[T]) Extreme(fn func(a, b T) bool) T {
	var result T
	for i, v := range s.data {
		if fn(v, result) || i == 0 {
			result = v
		}
	}
	return result
}

// Max 数据最大值, a > b
//
// # Examples:
//
// s := NewStream([]int{1, 2, 3, 4, 5})
//
// s.Max(func(a, b int) bool { return a > b }) // 5
func (s *Stream[T]) Max(fn func(a, b T) bool) T {
	return s.Extreme(fn)
}

// Min 数据最小值, a < b
//
// # Examples:
//
// s := NewStream([]int{1, 2, 3, 4, 5})
//
// s.Min(func(a, b int) bool { return a < b }) // 1
func (s *Stream[T]) Min(fn func(a, b T) bool) T {
	return s.Extreme(fn)
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
