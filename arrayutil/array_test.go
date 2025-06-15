package arrayutil

import (
	"testing"
)

func TestDuplicate(t *testing.T) {
	t.Log(Union([]int{1, 2, 3}, []int{2, 3, 4})) // [1 2 3 4]
}

func TestContain(t *testing.T) {
	t.Log(Contains([]int{12, 13, 1}, 12))            // true
	t.Log(Contains([]string{"12", "13", "1"}, "12")) // true
}

func TestJoin(t *testing.T) {
	t.Log(Join([]int8{12, 22, 12}, ","))           // 12,22,12
	t.Log(Join([]int16{12, 22, 12}, ","))          // 12,22,12
	t.Log(Join([]int{12, 22, 12}, ","))            // 12,22,12
	t.Log(Join([]uint{12, 22, 12}, ","))           // 12,22,12
	t.Log(Join([]float32{12.3, 12.31, 12}, ","))   // 12.3,12.31,12
	t.Log(Join([]float64{12.3, 12.31, 12}, ","))   // 12.3,12.31,12
	t.Log(Join([]string{"12", "321", "000"}, ",")) // 12,321,000
}

func TestConcat(t *testing.T) {
	t.Log(Concat([]int{12, 13, 1}, []int{22, 12, 1}))                   // [12 13 1 22 12 1]
	t.Log(Concat([]string{"12", "13", "1"}, []string{"22", "12", "1"})) // [12 13 1 22 12 1]
}

func TestIntersect(t *testing.T) {
	t.Log(Intersect([]int{12, 13, 1}, []int{22, 12, 1}))                    // [12 1]
	t.Log(Intersect([]int{12, 13, 1}, []int{22, 12, 1}, []int{12, 13, 14})) // [12]
	t.Log(Intersect([]string{"12", "13", "1"}, []string{"22", "12", "1"}))  // [12 1]
	t.Log(Intersect[int]())                                                 // []
	t.Log(Intersect([]int{12, 22, 12}, []int{12, 22}))                      // [12 22]
}
