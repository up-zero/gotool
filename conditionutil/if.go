package conditionutil

// If 三元运算符
//
// # Params:
//
//	condition: 条件
//	a: 条件为真时返回的值
//	b: 条件为假时返回的值
func If(condition bool, a, b any) any {
	if condition {
		return a
	}
	return b
}
