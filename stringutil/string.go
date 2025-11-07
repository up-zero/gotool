package stringutil

import (
	"strings"
)

// Reverse 字符串反转
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// TakeFirst 截取字符串前 n 个 Unicode 字符（rune）
//
// # Params:
//
//	s: 字符串
//	maxLength: 字符长度
//
// # Examples:
//
//	Truncate("hello world", 5) // hello
//	Truncate("你好 世界!", 2) // 你好
func TakeFirst(s string, n int) string {
	if n <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}

// ContainsAny 判断字符串 s 是否包含 substrs 中的任意一个子串
//
// # Examples:
//
//	ContainsAny("hello world", "world", "123") // true
//	ContainsAny("hello world", "123") // false
func ContainsAny(s string, substrs ...string) bool {
	for _, sub := range substrs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// ContainsAll 判断字符串 s 是否包含 substrs 中的所有子串
//
// # Examples:
//
//	ContainsAll("hello world", "hello", "world") // true
//	ContainsAll("hello world", "hello", "world", "123") // false
func ContainsAll(s string, substrs ...string) bool {
	for _, sub := range substrs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}
