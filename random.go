package gotool

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomStr 随机字符串
//
// str: 待随机的字符串
// length: 随机字符串长度
func RandomStr(str string, length int) string {
	if length <= 0 {
		return ""
	}
	var ans string
	for i := 0; i < length; i++ {
		ans += string(str[r.Intn(len(str))])
	}
	return ans
}

// RandomNumber 随机数字
//
// length: 随机数字长度
func RandomNumber(length int) string {
	return RandomStr("0123456789", length)
}

// RandomAlpha 随机字母
//
// length: 随机字母长度
func RandomAlpha(length int) string {
	return RandomStr("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz", length)
}

// RandomAlphaNumber 随机字母数字
//
// length: 随机字母数字长度
func RandomAlphaNumber(length int) string {
	return RandomStr("0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz", length)
}

// RandomRangeInt 指定范围内的随机数 [最小值, 最大值)
//
// minValue 最小值（包含）
// maxValue 最大值（不包含）
func RandomRangeInt(minValue, maxValue int) int {
	if minValue >= maxValue {
		return 0
	}
	return r.Intn(maxValue-minValue) + minValue
}
