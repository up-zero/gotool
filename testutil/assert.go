package testutil

import (
	"math"
	"reflect"
	"testing"
)

// defaultEpsilon 默认的误差范围
const defaultEpsilon = 1e-5

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

// EqualFloat 断言两个值（浮点数或浮点数切片）在误差范围内相等
//
// # Params:
//
//	t: 测试对象
//	got: 实际值
//	want: 期望值
//	epsilons: 允许的误差范围 (例如 0.00001)
func EqualFloat(t *testing.T, got, want any, epsilons ...float64) {
	t.Helper()

	gotVal := reflect.ValueOf(got)
	wantVal := reflect.ValueOf(want)

	// 类型判断
	if gotVal.Type() != wantVal.Type() {
		t.Fatalf("assertion failed: types mismatch\n"+
			"got type:  %T\n"+
			"want type: %T", got, want)
		return
	}

	epsilon := defaultEpsilon
	if len(epsilons) > 0 {
		epsilon = epsilons[0]
	}

	switch gotVal.Kind() {
	case reflect.Float32, reflect.Float64:
		// 处理单个浮点数
		g := gotVal.Float()
		w := wantVal.Float()
		if math.Abs(g-w) > epsilon {
			t.Helper()
			t.Fatalf("assertion failed, float mismatch within epsilon %v:\n"+
				"got:  %v\n"+
				"want: %v", epsilon, got, want)
		}

	case reflect.Slice, reflect.Array:
		// 处理数组或切片
		if gotVal.Len() != wantVal.Len() {
			t.Fatalf("assertion failed: slice lengths differ\n"+
				"got len:  %d\n"+
				"want len: %d", gotVal.Len(), wantVal.Len())
		}

		// 遍历每个元素
		for i := 0; i < gotVal.Len(); i++ {
			gEle := gotVal.Index(i)
			wEle := wantVal.Index(i)

			// 检查元素是否为浮点数
			if !isFloat(gEle.Kind()) {
				t.Fatalf("EqualFloat only supports floats, found %v at index %d", gEle.Kind(), i)
			}

			if math.Abs(gEle.Float()-wEle.Float()) > epsilon {
				t.Fatalf("assertion failed at index %d:\n"+
					"got:     %v\n"+
					"want:    %v\n"+
					"epsilon: %v", i, gEle.Float(), wEle.Float(), epsilon)
			}
		}

	default:
		t.Fatalf("EqualFloat only supports float or slice of floats, but got type: %T", got)
	}
}

// isFloat 检查给定的 reflect.Kind 是否为浮点数
func isFloat(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
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
