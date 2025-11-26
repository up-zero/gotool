package convertutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestCopyProperties(t *testing.T) {
	type src struct {
		Name string
		Map  map[string]int
	}
	type dst struct {
		Name string
		Map  map[string]int
		Age  int
	}
	// 1. normal struct
	s1 := src{Name: "test", Map: map[string]int{"a": 1}}
	s2 := new(dst)
	if err := CopyProperties(s1, s2); err != nil {
		t.Fatal(err)
	}
	// &{test map[a:1] 0}
	testutil.Equal(t, s2, &dst{Name: "test", Map: map[string]int{"a": 1}, Age: 0})

	// 2. ptr struct
	s3 := &src{Name: "test", Map: map[string]int{"a": 3}}
	s4 := new(dst)
	if err := CopyProperties(s3, s4); err != nil {
		t.Fatal(err)
	}
	// &{test map[a:3] 0}
	testutil.Equal(t, s4, &dst{Name: "test", Map: map[string]int{"a": 3}, Age: 0})
}
