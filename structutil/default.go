package structutil

import (
	"github.com/up-zero/gotool"
	"reflect"
	"strconv"
)

// SetDefaults 设置结构体默认值，根据结构体标签 "default" 为指针指向的结构体实例设置默认值。
//
// # Params:
//
// ptr: 指向结构体的指针
//
// # Examples:
//
//	type A struct {
//		A string `default:"a"`
//		B int
//	}
//	type d struct {
//		F string `default:"f"`
//		*A
//	}
//
//	dd := new(d)
//	if err := SetDefaults(dd); err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("%+v \n", dd) // &{F:f A:0xc0000082d0}
//	log.Printf("%+v \n", dd.A) // &{A:a B:0}
func SetDefaults(ptr any) error {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return gotool.ErrDstMustBePointerStruct
	}

	return setDefaultsRecursive(v.Elem())
}

// setDefaultsRecursive 递归地为结构体及其嵌套结构体设置默认值
func setDefaultsRecursive(elem reflect.Value) error {
	t := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldType := t.Field(i)

		if !field.CanSet() {
			continue
		}

		// 如果字段是结构体，递归处理
		if field.Kind() == reflect.Struct {
			if err := setDefaultsRecursive(field); err != nil {
				return err
			}
			continue
		}

		// 如果字段是结构体的指针，检查是否为 nil 并初始化
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			if field.IsNil() {
				// 初始化指针指向的结构体
				field.Set(reflect.New(field.Type().Elem()))
			}
			// 递归处理指针指向的结构体
			if err := setDefaultsRecursive(field.Elem()); err != nil {
				return err
			}
			continue
		}

		if field.IsZero() {
			defaultValue := fieldType.Tag.Get("default")
			if defaultValue == "" {
				continue
			}
			if err := setField(field, defaultValue); err != nil {
				return err
			}
		}
	}
	return nil
}

func setField(field reflect.Value, defaultValue string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(defaultValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(defaultValue, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(defaultValue, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(val)
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(defaultValue, 64)
		if err != nil {
			return err
		}
		field.SetFloat(val)
	case reflect.Bool:
		val, err := strconv.ParseBool(defaultValue)
		if err != nil {
			return err
		}
		field.SetBool(val)
	default:
		return gotool.ErrNotSupportFormat
	}
	return nil
}
