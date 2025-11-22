package stringutil

import (
	"strings"
	"unicode"
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

// CamelToSnake 将驼峰式字符串转换为下划线连接
//
//   - PascalCase -> pascal_case
//   - lowerCamelCase -> lower_camel_case
//   - HTTPRequest -> http_request
//   - MyID -> my_id
//   - ID -> id
func CamelToSnake(s string) string {
	var builder strings.Builder
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}

	for i, r := range runes {
		if unicode.IsUpper(r) {
			// 如果不是第一个字符，需要判断是否添加下划线
			if i > 0 {
				prev := runes[i-1]
				// 1. 前一个字符是小写或数字 (e.g., "myN", "v1G")
				// 2. 前一个是大写，但后一个(如果存在)是小写 (e.g., "HTTPRequest" 中的 'R')
				nextExists := i+1 < len(runes)
				if unicode.IsLower(prev) || unicode.IsDigit(prev) ||
					(unicode.IsUpper(prev) && nextExists && unicode.IsLower(runes[i+1])) {
					builder.WriteRune('_')
				}
			}
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// SnakeToCamel 将下划线连接的字符串转换为小驼峰式
//
//   - my_var -> myVar
//   - http_request -> httpRequest
//   - _my_var -> myVar
func SnakeToCamel(s string) string {
	var builder strings.Builder
	parts := strings.Split(s, "_")

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		if builder.Len() == 0 {
			// 首字母小写
			builder.WriteString(part)
		} else {
			// 后序首字母大写
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			builder.WriteString(string(runes))
		}
	}
	return builder.String()
}

// SnakeToPascal 将下划线连接的字符串转换为大驼峰式（PascalCase）
//
//   - my_var -> MyVar
//   - http_request -> HttpRequest
//   - _my_var -> MyVar
func SnakeToPascal(s string) string {
	var builder strings.Builder
	parts := strings.Split(s, "_")

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		builder.WriteString(string(runes))
	}
	return builder.String()
}
