package testutil

import (
	"reflect"
	"testing"
)

// Equal 断言两个值相等
//
// # Params:
//
//	t: 测试对象
//	got: 实际值
//	want: 期望值
func Equal(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("assertion failed, values not equal: \n"+
			"got:  %#v (type: %T)\n"+
			"want: %#v (type: %T)", got, got, want, want)
	}
}

// NotEqual 断言两个值不相等
//
// # Params:
//
//	t: 测试对象
//	got: 实际值
//	want: 期望值
func NotEqual(t *testing.T, got, want any) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Fatalf("assertion failed, values is equal: \n"+
			"got:  %#v (type: %T)", got, got)
	}
}
