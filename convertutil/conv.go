package convertutil

import (
	"fmt"
	"strconv"
)

// StrToInt8 字符串转换为int8
func StrToInt8(s string) int8 {
	num, _ := strconv.ParseInt(s, 10, 8)
	return int8(num)
}

// StrToInt16 字符串转换为int16
func StrToInt16(s string) int16 {
	num, _ := strconv.ParseInt(s, 10, 16)
	return int16(num)
}

// StrToInt32 字符串转换为int32
func StrToInt32(s string) int32 {
	num, _ := strconv.ParseInt(s, 10, 32)
	return int32(num)
}

// StrToInt64 字符串转换为int64
func StrToInt64(s string) int64 {
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}

// StrToUint8 字符串转换为uint8
func StrToUint8(s string) uint8 {
	num, _ := strconv.ParseUint(s, 10, 8)
	return uint8(num)
}

// StrToUint16 字符串转换为uint16
func StrToUint16(s string) uint16 {
	num, _ := strconv.ParseUint(s, 10, 16)
	return uint16(num)
}

// StrToUint32 字符串转换为uint32
func StrToUint32(s string) uint32 {
	num, _ := strconv.ParseUint(s, 10, 32)
	return uint32(num)
}

// StrToUint64 字符串转换为uint64
func StrToUint64(s string) uint64 {
	num, _ := strconv.ParseUint(s, 10, 64)
	return num
}

// StrToFloat32 字符串转换为float32
func StrToFloat32(s string) float32 {
	num, _ := strconv.ParseFloat(s, 32)
	return float32(num)
}

// StrToFloat64 字符串转换为float64
func StrToFloat64(s string) float64 {
	num, _ := strconv.ParseFloat(s, 64)
	return num
}

// Int64ToStr int64转换为字符串
func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Uint64ToStr uint64转换为字符串
func Uint64ToStr(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// Float64ToStr float64转换为字符串
//
// # Params
//
//	num: 待转化的数字
//	prec: 精度, 及就是小数点后的位数, 默认为-1, 即不限制
func Float64ToStr(num float64, prec ...int) string {
	if len(prec) > 0 {
		return strconv.FormatFloat(num, 'f', prec[0], 64)
	}
	return strconv.FormatFloat(num, 'f', -1, 64)
}

// Int64ToHex int64转换为十六进制字符串
//
// # Params
//
//	num: int64
//	args: 可选参数, 用于指定填充的位数以及填充的字
//
// # Examples:
//
//	Int64ToHex(15) // 返回 F
//	Int64ToHex(15, "08") // 返回 0000000F
func Int64ToHex(num int64, args ...string) string {
	var format string
	if len(args) > 0 {
		format = args[0]
	}
	return fmt.Sprintf("%"+format+"X", num)
}

// HexToInt64 十六进制字符串转换为int64
func HexToInt64(str string) int64 {
	num, _ := strconv.ParseInt(str, 16, 64)
	return num
}
