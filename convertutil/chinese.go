package convertutil

import (
	"strconv"
	"strings"
)

// 数字汉字映射表
var digitToChineseMap = map[rune]string{
	'0': "零", '1': "一", '2': "二", '3': "三", '4': "四",
	'5': "五", '6': "六", '7': "七", '8': "八", '9': "九",
}

// DigitToChinese 数字逐位转汉字
//
// # Params
//
//	s: 待转换的数字，例如：138000
//
// # Examples:
//
//	DigitToChinese("138000") // 返回 一三八零零零
func DigitToChinese(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(digitToChineseMap[r])
	}
	return sb.String()
}

// IntegerToChinese 整数转中文读法
//
// # Params
//
//	num: 待转换的整数
//
// # Examples:
//
//	IntegerToChinese(138000) // 输出 十三万八千
//	IntegerToChinese(-10001)   // 输出 负一万零一
func IntegerToChinese(num int) string {
	return integerToChineseInternal(num, true)
}

// integerToChineseInternal 整数转换为中文读法内部递归函数
//
// # Params
//
//	num: 待转换的数字
//	isTopLevel: 是否处于数字最高位（控制十位数字1的省略逻辑）
func integerToChineseInternal(num int, isTopLevel bool) string {
	if num == 0 {
		if isTopLevel {
			return "零" // 仅当0是整个数字时返回"零"
		}
		return "" // 中间0由调用方处理
	}
	if num < 0 {
		return "负" + integerToChineseInternal(-num, true)
	}

	// 处理亿级 (10^8)
	if num >= 100000000 {
		high := num / 100000000
		low := num % 100000000
		highStr := integerToChineseInternal(high, isTopLevel)

		if low == 0 {
			return highStr + "亿"
		}

		lowStr := integerToChineseInternal(low, false)
		if low < 10000000 { // 需要补零的情况 (1,0000,0001 -> 一亿零一)
			return highStr + "亿零" + lowStr
		}
		return highStr + "亿" + lowStr
	}

	// 处理万级 (10^4)
	if num >= 10000 {
		high := num / 10000
		low := num % 10000
		highStr := integerToChineseInternal(high, isTopLevel)

		if low == 0 {
			return highStr + "万"
		}

		lowStr := integerToChineseInternal(low, false)
		if low < 1000 { // 需要补零的情况 (1,0001 -> 一万零一)
			return highStr + "万零" + lowStr
		}
		return highStr + "万" + lowStr
	}

	// 处理0-9999
	return integerToChineseSub10k(num, isTopLevel)
}

// integerToChineseSub10k 处理0-9999的数字
//
// # Params
//
//	num: 待转换的数字
//	isTopLevel: 是否处于数字最高位（控制十位数字1的省略逻辑）
func integerToChineseSub10k(num int, isTopLevel bool) string {
	if num == 0 {
		return ""
	}
	s := strconv.Itoa(num)
	// 预定义单位（千、百、十、个）
	unitArray := []string{"千", "百", "十", ""}
	startIndex := 4 - len(s)
	units := unitArray[startIndex : startIndex+len(s)]
	digits := []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九"}

	var result strings.Builder
	zero := false // 标记是否需要添加"零"

	for i, char := range s {
		digit := int(char - '0')
		if digit == 0 {
			// 标记需要补零（后续遇到非零数字时添加）
			if i < len(s)-1 {
				nextDigit := int(s[i+1] - '0')
				if nextDigit != 0 {
					zero = true
				}
			}
			continue
		}

		// 需要补零时添加
		if zero {
			result.WriteString("零")
			zero = false
		}

		// 十位数字1的特殊处理
		current := digits[digit]
		unit := units[i]

		// 规则1: 10-19在最高位时省略"一"（如10->"十"）
		// 规则2: 非最高位或非10-19时保留"一"（如10010中的"一十"）
		if digit == 1 && unit == "十" {
			if isTopLevel && len(s) == 2 && i == 0 {
				// 10-19且处于最高位：省略"一"
				current = ""
			}
		}

		result.WriteString(current + unit)
	}

	return result.String()
}
