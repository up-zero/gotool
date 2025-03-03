package gotool

import "testing"

func TestArrayDuplicate(t *testing.T) {
	t.Log(ArrayDuplicate([]int{1, 2, 3}, []int{2, 3, 4})) // [1 2 3 4]
}

func TestArrayIn(t *testing.T) {
	t.Log(ArrayIn(12, []int{12, 13, 1})) // true
}

func TestArrayJoin(t *testing.T) {
	t.Log(ArrayJoin([]int{12, 22, 12}, ",")) // 12,22,12
}
