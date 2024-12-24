package gotool

import (
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
	t.Log(s2)

	// 2. ptr struct
	s3 := &src{Name: "test", Map: map[string]int{"a": 3}}
	s4 := new(dst)
	if err := CopyProperties(s3, s4); err != nil {
		t.Fatal(err)
	}
	t.Log(s4)
}
