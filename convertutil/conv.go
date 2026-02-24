package convertutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// StrToInt 字符串转换为int
func StrToInt(s string) int {
	num, _ := strconv.ParseInt(s, 10, 64)
	return int(num)
}

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

// StrToUint 字符串转换为uint
func StrToUint(s string) uint {
	num, _ := strconv.ParseUint(s, 10, 64)
	return uint(num)
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

// ToStr 将任意类型的值转换为字符串
//
// # Params
//
//	value: 待转换的值
//
// # Examples:
//
//	ToStr(123) // 123
//	ToStr("123") // 123
//	ToStr(123.456) // 123.456
//	ToStr([]int{1, 2, 3}) // [1,2,3]
func ToStr(value any) string {
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case bool:
		return strconv.FormatBool(val)
	case []byte:
		return string(val)
	case error:
		return val.Error()
	case fmt.Stringer:
		return val.String()
	}

	rv := reflect.ValueOf(value)

	// 递归解指针
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return ""
		}
		return ToStr(rv.Elem().Interface())
	}

	// 处理自定义别名类型
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	case reflect.String:
		return rv.String()
	}

	// 对结构体、Map、Slice 等序列化
	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b)
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
