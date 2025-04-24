package mathutil

import (
	"fmt"
	"testing"
)

func TestMathAbs(t *testing.T) {
	fmt.Println(MathAbs(-42))        // 输出: 42
	fmt.Println(MathAbs(-123456789)) // 输出: 123456789

	// 测试浮点数
	fmt.Println(MathAbs(-3.14)) // 输出: 3.14
	fmt.Println(MathAbs(-0.0))  // 输出: 0.0
}

func TestMathMin(t *testing.T) {
	fmt.Println(MathMin(12, 20, -1))   // 输出: -1
	fmt.Println(MathMin(12, 20, -1.1)) // 输出: -1.1
}

func TestMathMax(t *testing.T) {
	fmt.Println(MathMax(12, 200, -1))       // 输出: 200
	fmt.Println(MathMax(-100, -1.12, -1.1)) // 输出: -1.1
}
