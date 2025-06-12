package arrayutil

import (
	"testing"
)

func TestDuplicate(t *testing.T) {
	t.Log(Duplicate([]int{1, 2, 3}, []int{2, 3, 4})) // [1 2 3 4]
}

func TestContain(t *testing.T) {
	t.Log(Contains([]int{12, 13, 1}, 12))            // true
	t.Log(Contains([]string{"12", "13", "1"}, "12")) // true
}

func TestArrayJoin(t *testing.T) {
	t.Log(Join([]int8{12, 22, 12}, ","))           // 12,22,12
	t.Log(Join([]int16{12, 22, 12}, ","))          // 12,22,12
	t.Log(Join([]int{12, 22, 12}, ","))            // 12,22,12
	t.Log(Join([]uint{12, 22, 12}, ","))           // 12,22,12
	t.Log(Join([]float32{12.3, 12.31, 12}, ","))   // 12.3,12.31,12
	t.Log(Join([]float64{12.3, 12.31, 12}, ","))   // 12.3,12.31,12
	t.Log(Join([]string{"12", "321", "000"}, ",")) // 12,321,000
}
