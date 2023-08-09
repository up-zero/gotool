package gotool

import "testing"

func TestArrayDuplicate(t *testing.T) {
	t.Log(ArrayDuplicate([]int{1, 2, 3}, []int{2, 3, 4})) // [1 2 3 4]
}
