package validator

import (
	"net"
	"unicode"
)

// IsDigit 判断字符串是否为数字
func IsDigit(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsAlpha 验证字符串是否为 Unicode 字符
//
//   - a-z A-Z 这种为 Unicode 字符
//   - 汉字也是 Unicode 字符
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsAlphaStrict 验证字符串是否为英文字符 a-z A-Z
func IsAlphaStrict(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric 验证字符串是否为 Unicode 字符或数字
func IsAlphaNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsIpv4 验证字符串是否为 IPv4 地址
func IsIpv4(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	return ip.To4() != nil
}

// IsIpv6 验证字符串是否为 IPv6 地址
func IsIpv6(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	return ip.To4() == nil && ip.To16() != nil
}

// ContainChinese 判断字符串是否包含中文字符
func ContainChinese(s string) bool {
	for _, r := range s {
		if r >= 0x4e00 && r <= 0x9fa5 ||
			(r >= 0x3000 && r <= 0x303f) ||
			(r >= 0xff00 && r <= 0xffef) {
			return true
		}
	}
	return false
}
