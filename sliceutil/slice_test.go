package sliceutil

import (
	"fmt"
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestDuplicate(t *testing.T) {
	testutil.Equal(t, Unique([]int{1, 2, 3}, []int{2, 3, 4}), []int{1, 2, 3, 4})
}

func TestContain(t *testing.T) {
	testutil.Equal(t, Contains([]int{12, 13, 1}, 12), true)
	testutil.Equal(t, Contains([]string{"12", "13", "1"}, "12"), true)
}

func TestJoin(t *testing.T) {
	testutil.Equal(t, Join([]int8{12, 22, 12}, ","), "12,22,12")
	testutil.Equal(t, Join([]int16{12, 22, 12}, ","), "12,22,12")
	testutil.Equal(t, Join([]int{12, 22, 12}, ","), "12,22,12")
	testutil.Equal(t, Join([]uint{12, 22, 12}, ","), "12,22,12")
	testutil.Equal(t, Join([]float32{12.3, 12.31, 12}, ","), "12.3,12.31,12")
	testutil.Equal(t, Join([]float64{12.3, 12.31, 12}, ","), "12.3,12.31,12")
	testutil.Equal(t, Join([]string{"12", "321", "000"}, ","), "12,321,000")
}

func TestConcat(t *testing.T) {
	testutil.Equal(t, Concat([]int{12, 13, 1}, []int{22, 12, 1}), []int{12, 13, 1, 22, 12, 1})
	testutil.Equal(t, Concat([]string{"12", "13", "1"}, []string{"22", "12", "1"}), []string{"12", "13", "1", "22", "12", "1"})
}

func TestIntersect(t *testing.T) {
	testutil.Equal(t, Intersect([]int{12, 13, 1}, []int{22, 12, 1}), []int{12, 1})
	testutil.Equal(t, Intersect([]int{12, 13, 1}, []int{22, 12, 1}, []int{12, 13, 14}), []int{12})
	testutil.Equal(t, Intersect([]string{"12", "13", "1"}, []string{"22", "12", "1"}), []string{"12", "1"})
	testutil.Equal(t, Intersect[int](), []int{})
	testutil.Equal(t, Intersect([]int{12, 22, 12}, []int{12, 22}), []int{12, 22})

	t.Log(Intersect([]int{12, 13, 1}, []int{22, 12, 1}))                    // [12 1]
	t.Log(Intersect([]int{12, 13, 1}, []int{22, 12, 1}, []int{12, 13, 14})) // [12]
	t.Log(Intersect([]string{"12", "13", "1"}, []string{"22", "12", "1"}))  // [12 1]
	t.Log(Intersect[int]())                                                 // []
	t.Log(Intersect([]int{12, 22, 12}, []int{12, 22}))                      // [12 22]
}

func TestFilter(t *testing.T) {
	testutil.Equal(t, Filter([]int{12, 13, 1}, func(v int) bool { return v > 10 }), []int{12, 13})
}

func TestMap(t *testing.T) {
	testutil.Equal(t, Map([]int{12, 13, 1}, func(v int) string { return fmt.Sprintf("->%v", v) }), []string{"->12", "->13", "->1"})
}

func TestGroupBy(t *testing.T) {
	testutil.Equal(t, GroupBy([]int{12, 13, 1}, func(v int) string { return fmt.Sprintf("%v", v%2) }), map[string][]int{"0": {12}, "1": {13, 1}})
}
