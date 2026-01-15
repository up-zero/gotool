package randomutil

import (
	"math/rand"
	"strings"
)

// RandomStr 随机字符串
//
// # Params:
//
//	str: 待随机的字符串
//	length: 随机字符串长度
func RandomStr(str string, length int) string {
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

// RandomNumber 随机数字
//
// # Params:
//
//	length: 随机数字长度
func RandomNumber(length int) string {
	return RandomStr("0123456789", length)
}

// RandomAlpha 随机字母
//
// # Params:
//
//	length: 随机字母长度
func RandomAlpha(length int) string {
	return RandomStr("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz", length)
}

// RandomAlphaNumber 随机字母数字
//
// # Params:
//
//	length: 随机字母数字长度
func RandomAlphaNumber(length int) string {
	return RandomStr("0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz", length)
}

// RandomRangeInt 指定范围内的随机数 [最小值, 最大值)
//
// # Params:
//
//	minValue: 最小值（包含）
//	maxValue: 最大值（不包含）
func RandomRangeInt(minValue, maxValue int) int {
	if minValue >= maxValue {
		return 0
	}
	return rand.Intn(maxValue-minValue) + minValue
}
