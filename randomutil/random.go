package randomutil

import (
	"math/rand"
	"strings"
)

// String 随机字符串
//
// # Params:
//
//	str: 待随机的字符串
//	length: 随机字符串长度
func String(str string, length int) string {
	if length <= 0 {
		return ""
	}

	runes := []rune(str)
	var builder strings.Builder
	builder.Grow(length)

	for i := 0; i < length; i++ {
		builder.WriteRune(runes[rand.Intn(len(runes))])
	}
	return builder.String()
}

// Digits 随机数字
//
// # Params:
//
//	length: 随机数字长度
func Digits(length int) string {
	return String("0123456789", length)
}

// Letters 随机字母
//
// # Params:
//
//	length: 随机字母长度
func Letters(length int) string {
	return String("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz", length)
}

// Alphanumeric 随机字母数字
//
// # Params:
//
//	length: 随机字母数字长度
func Alphanumeric(length int) string {
	return String("0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz", length)
}

// IntRange 指定范围内的随机数 [最小值, 最大值)
//
// # Params:
//
//	minValue: 最小值（包含）
//	maxValue: 最大值（不包含）
func IntRange(minValue, maxValue int) int {
	if minValue >= maxValue {
		return 0
	}
	return rand.Intn(maxValue-minValue) + minValue
}

// Bool 随机布尔值
func Bool() bool {
	return rand.Float32() < 0.5
}

// Choice 从切片随机选一个元素
func Choice[T any](list []T) T {
	if len(list) == 0 {
		var zero T
		return zero
	}
	return list[rand.Intn(len(list))]
}
