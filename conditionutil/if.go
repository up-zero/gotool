package conditionutil

// If 三元运算符
//
// # Params:
//
//	condition: 条件
//	a: 条件为真时返回的值
//	b: 条件为假时返回的值
func If[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// IfLazy 延迟执行的三元运算
//
// # Params:
//
//	condition: 条件
//	a: 条件为真时调用的函数
//	b: 条件为假时调用的函数
func IfLazy[T any](condition bool, a, b func() T) T {
	if condition {
		return a()
	}
	return b()
}
