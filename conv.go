package gotool

import "strconv"

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
