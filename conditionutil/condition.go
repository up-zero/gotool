package conditionutil

import "reflect"

// IsZero 判断一个值是否为零值
//
// # Examples:
//
//	IsZero(User{}) // true
//	IsZero(User{Name: "test"}) // false
func IsZero(v any) bool {
	if v == nil {
		return true
	}
	val := reflect.ValueOf(v)

	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return true
		}
		val = val.Elem()
	}

	return val.IsZero()
}
