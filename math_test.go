package gotool

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
