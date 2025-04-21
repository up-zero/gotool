package gotool

import (
	"fmt"
	"testing"
)

func TestStream_Filter(t *testing.T) {
	s := NewStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// []int{2, 4, 6, 8, 10}
	fmt.Printf("%#v \n", s.Filter(func(v int) bool { return v%2 == 0 }).ToSlice())
}

func TestStream_Map(t *testing.T) {
	s := NewStream([]int{1, 2, 3, 4, 5})
	s.Map(func(v int) int { return v * 2 })
	// []int{2, 4, 6, 8, 10}
	fmt.Printf("%#v \n", s.ToSlice())
}

func TestStreamMap(t *testing.T) {
	s := NewStream([]int{1, 2, 3, 4, 5})
	s2 := StreamMap(s, func(v int) string { return fmt.Sprintf("data: %d", v) })
	// []string{"data: 1", "data: 2", "data: 3", "data: 4", "data: 5"}
	fmt.Printf("%#v \n", s2.ToSlice())
}

func TestStream_Extreme(t *testing.T) {
	s := NewStream([]int{1, 2, 3, 4, 5})
	println(s.Extreme(func(a, b int) bool { return a > b })) // 5
	println(s.Extreme(func(a, b int) bool { return a < b })) // 1
}

func TestStream_Max(t *testing.T) {
	s := NewStream([]int{1, 2, 3, 4, 5})
	println(s.Max(func(a, b int) bool { return a > b })) // 5
}

func TestStream_Min(t *testing.T) {
	s := NewStream([]int{1, 2, 3, 4, 5})
	println(s.Min(func(a, b int) bool { return a < b })) // 1
}
