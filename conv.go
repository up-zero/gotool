package gotool

import (
	"fmt"
	"strconv"
)

// StrToInt64 字符串转换为int64
//
// str: 字符串
func StrToInt64(str string) int64 {
	num, _ := strconv.ParseInt(str, 10, 64)
	return num
}

// StrToUint64 字符串转换为uint64
//
// str: 字符串
func StrToUint64(str string) uint64 {
	num, _ := strconv.ParseUint(str, 10, 64)
	return num
}

// StrToFloat64 字符串转换为float64
//
// str: 字符串
func StrToFloat64(str string) float64 {
	num, _ := strconv.ParseFloat(str, 64)
	return num
}

// Int64ToStr int64转换为字符串
//
// num: int64
func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Uint64ToStr uint64转换为字符串
//
// num: uint64
func Uint64ToStr(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// Float64ToStr float64转换为字符串
//
// num: float64
// prec: 精度, 及就是小数点后的位数, 默认为-1, 即不限制
func Float64ToStr(num float64, prec ...int) string {
	if len(prec) > 0 {
		return strconv.FormatFloat(num, 'f', prec[0], 64)
	}
	return strconv.FormatFloat(num, 'f', -1, 64)
}

// Int64ToHex int64转换为十六进制字符串
//
// num: int64
// args: 可选参数, 用于指定填充的位数以及填充的字
//
// Examples:
//	gotool.Int64ToHex(15) // 返回 F
//	gotool.Int64ToHex(15, "08") // 返回 000F
func Int64ToHex(num int64, args ...string) string {
	var format string
	if len(args) > 0 {
		format = args[0]
	}
	return fmt.Sprintf("%"+format+"X", num)
}

// HexToInt64 十六进制字符串转换为int64
//
// str: 十六进制字符串
func HexToInt64(str string) int64 {
	num, _ := strconv.ParseInt(str, 16, 64)
	return num
}
