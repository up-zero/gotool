package mathutil

import (
	"fmt"
	"testing"
)

func TestMathAbs(t *testing.T) {
	fmt.Println(Abs(-42))        // 输出: 42
	fmt.Println(Abs(-123456789)) // 输出: 123456789

	// 测试浮点数
	fmt.Println(Abs(-3.14)) // 输出: 3.14
	fmt.Println(Abs(-0.0))  // 输出: 0.0
}

func TestMathMin(t *testing.T) {
	fmt.Println(Min(12, 20, -1))   // 输出: -1
	fmt.Println(Min(12, 20, -1.1)) // 输出: -1.1
}

func TestMathMax(t *testing.T) {
	fmt.Println(Max(12, 200, -1))       // 输出: 200
	fmt.Println(Max(-100, -1.12, -1.1)) // 输出: -1.1
}
