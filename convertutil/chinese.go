package convertutil

import (
	"regexp"
	"strconv"
	"strings"
)

// 数字汉字映射表
var digitToChineseMap = map[rune]string{
	'0': "零", '1': "一", '2': "二", '3': "三", '4': "四",
	'5': "五", '6': "六", '7': "七", '8': "八", '9': "九",
}

// 预编译正则 (Regex)
var (
	// [Time] 匹配时间：12:30, 09:05
	reTime = regexp.MustCompile(`(\d{1,2}):(\d{2})`)

	// [Date] 匹配年份：2025年
	reDateYear = regexp.MustCompile(`(\d{4})年`)

	// [Date] 匹配日期：5月20日
	reDateDay = regexp.MustCompile(`(\d{1,2})月(\d{1,2})日`)

	// [Percentage] 匹配百分比：50%, 3.5%, 100%
	// 包含整数和浮点数的情况
	rePercent = regexp.MustCompile(`(\d+(?:\.\d+)?)%`)

	// [Decimal] 匹配浮点数：3.14, 10.5
	reDecimal = regexp.MustCompile(`(\d+)\.(\d+)`)

	// [Phone] 匹配手机号：11位数字
	rePhone = regexp.MustCompile(`\d{11}`)

	// [Number] 匹配纯数字
	reNum = regexp.MustCompile(`\d+`)
)

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

// punctuationReplacer 标点替换器
var punctuationReplacer = strings.NewReplacer(
	"，", ",",
	"。", ".",
	"！", "!",
	"？", "?",
	"；", ",",
	"：", ",", // 中文冒号转逗号
	"、", ",",
)

// TextToChinese 中文文本口语化转换
//
// # Params:
//
//	text: 待处理的文本
//
// # Examples:
//
//	TextToChinese("12:30") // 十二点三十分
//	TextToChinese("50%") // 百分之五十
//	TextToChinese("3.14") // 三点一四
func TextToChinese(text string) string {
	// processDecimal 处理浮点数读法: 3.14 -> 三点一四
	processDecimal := func(s string) string {
		parts := strings.Split(s, ".")
		if len(parts) != 2 {
			return s
		}
		// 整数部分：按数值读 (12.5 -> 十二点...)
		intPart, _ := strconv.Atoi(parts[0])
		intText := IntegerToChinese(intPart)

		// 小数部分：按位读
		decText := DigitToChinese(parts[1])

		return intText + "点" + decText
	}

	// processInteger 处理整数字符串
	processInteger := func(s string) string {
		// 太长的数字(>12位)按位读，否则按数值读
		if len(s) > 12 {
			return DigitToChinese(s)
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return DigitToChinese(s) // 溢出兜底
		}
		return IntegerToChinese(n)
	}

	// [Time] 时间处理: 12:30 -> 十二点三十分
	text = reTime.ReplaceAllStringFunc(text, func(s string) string {
		matches := reTime.FindStringSubmatch(s)
		h, _ := strconv.Atoi(matches[1])
		m, _ := strconv.Atoi(matches[2])
		// 读法：十二点三十分 (分钟为0时，如 12:00 -> 十二点整)
		if m == 0 {
			return IntegerToChinese(h) + "点整"
		}
		// 分钟数处理：05 -> 零五
		mStr := IntegerToChinese(m)
		if m < 10 && len(matches[2]) == 2 {
			mStr = "零" + mStr
		}
		return IntegerToChinese(h) + "点" + mStr + "分"
	})

	// [Percentage] 百分比处理: 50% -> 百分之五十
	text = rePercent.ReplaceAllStringFunc(text, func(s string) string {
		matches := rePercent.FindStringSubmatch(s)
		numberPart := matches[1] // 获取数字部分，可能是 "50" 或 "3.5"

		// 小数
		if strings.Contains(numberPart, ".") {
			return "百分之" + processDecimal(numberPart)
		}
		return "百分之" + processInteger(numberPart)
	})

	// [Date] 年份: 2025年 -> 二零二五年
	text = reDateYear.ReplaceAllStringFunc(text, func(s string) string {
		return DigitToChinese(s[:4]) + "年"
	})

	// [Date] 日期: 5月20日 -> 五月二十日
	text = reDateDay.ReplaceAllStringFunc(text, func(s string) string {
		matches := reDateDay.FindStringSubmatch(s)
		m, _ := strconv.Atoi(matches[1])
		d, _ := strconv.Atoi(matches[2])
		return IntegerToChinese(m) + "月" + IntegerToChinese(d) + "日"
	})

	// [Decimal] 浮点数: 3.14 -> 三点一四
	text = reDecimal.ReplaceAllStringFunc(text, func(s string) string {
		return processDecimal(s)
	})

	// [Phone] 手机号: 11位 -> 按位读
	text = rePhone.ReplaceAllStringFunc(text, func(s string) string {
		return DigitToChinese(s)
	})

	// [Integer] 常规整数
	text = reNum.ReplaceAllStringFunc(text, func(s string) string {
		return processInteger(s)
	})

	// 将中文标点转为英文
	text = punctuationReplacer.Replace(text)

	text = strings.ReplaceAll(text, ":", ",")

	return text
}
