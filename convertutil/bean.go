package convertutil

import (
	"reflect"

	"github.com/up-zero/gotool"
)

// CopyProperties 复制结构体的属性，注意：dst必须是指针类型，且指向的结构体类型与src类型相同
//
// # Params:
//
//	src: 源对象
//	dst: 目标对象
//
// # Examples:
//
//	type src struct {
//		Name string
//		Map  map[string]int
//	}
//	type dst struct {
//		Name string
//		Map  map[string]int
//		Age  int
//	}
//	s1 := src{Name: "test", Map: map[string]int{"a": 1}}
//	s2 := new(dst)
//	CopyProperties(s1, s2)
func CopyProperties(src, dst any) error {
	if src == nil || dst == nil {
		return gotool.ErrSrcDstCannotBeNil
	}

	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}
	dstValue := reflect.ValueOf(dst)

	if dstValue.Kind() != reflect.Ptr || dstValue.Elem().Kind() != reflect.Struct {
		return gotool.ErrDstMustBePointerStruct
	}
	dstValue = dstValue.Elem()

	return copyPropertiesRecursive(srcValue, dstValue)
}

func copyPropertiesRecursive(srcVal, dstVal reflect.Value) error {
	srcType := srcVal.Type()

	for i := 0; i < srcVal.NumField(); i++ {
		srcFieldType := srcType.Field(i)
		srcFieldValue := srcVal.Field(i)

		// 检查字段是否可导出
		if !srcFieldValue.CanInterface() {
			continue
		}

		dstField := dstVal.FieldByName(srcFieldType.Name)
		if dstField.IsValid() && dstField.CanSet() && dstField.Type() == srcFieldValue.Type() {
			dstField.Set(srcFieldValue)
		}

		// 嵌入结构体
		if srcFieldType.Anonymous {
			resolveValue := srcFieldValue
			if resolveValue.Kind() == reflect.Ptr {
				if resolveValue.IsNil() {
					continue
				}
				resolveValue = resolveValue.Elem()
			}

			if resolveValue.Kind() == reflect.Struct {
				if err := copyPropertiesRecursive(resolveValue, dstVal); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
