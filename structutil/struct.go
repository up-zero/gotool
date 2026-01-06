package structutil

import (
	"github.com/up-zero/gotool"
	"reflect"
	"strings"
)

// ToMap 将结构体转换为 map[string]any
//
// # Params:
//
//	in: 结构体实例
//
// # 特性：
//  1. 支持嵌套结构体（递归转换）
//  2. 支持匿名字段/嵌入结构体（自动展开，类似 json 行为）
//  3. 支持 json 标签（"name", "omitempty", "-"）
//  4. 保持原生类型
//
// # Examples:
//
//	type Base struct {
//		ID int `json:"id"`
//	}
//	type Info struct {
//		Age int `json:"age"`
//	}
//	type User struct {
//		Base
//		Name string `json:"name"`
//		Info Info   `json:"info"`
//	}
//	u := User{Base: Base{ID: 1}, Name: "Tom", Info: Info{Age: 18}}
//	m, _ := ToMap(u)
//	// m: {"id": 1, "name": "Tom", "info": {"age": 18}}
func ToMap(in any) (map[string]any, error) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, gotool.ErrMustBeStruct
	}

	return structToMap(v), nil
}

// structToMap 递归处理结构体转 map
func structToMap(v reflect.Value) map[string]any {
	out := make(map[string]any)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		fieldVal := v.Field(i)

		if !fieldVal.CanInterface() {
			continue
		}

		// json tag
		tag := fieldType.Tag.Get("json")
		if tag == "-" {
			continue
		}

		name, opts, _ := strings.Cut(tag, ",")

		// omitempty
		fieldVal.IsZero()
		if strings.Contains(opts, "omitempty") && fieldVal.IsZero() {
			continue
		}

		realVal := fieldVal
		if realVal.Kind() == reflect.Ptr {
			realVal = realVal.Elem()
		}

		// 匿名结构体展开
		if fieldType.Anonymous && name == "" {
			if realVal.Kind() == reflect.Struct {
				embeddedMap := structToMap(realVal)
				// 将匿名字段的 map 合并到当前 map 中
				// 当前 map 已有同名 key，保持当前 map 的（外部覆盖内部）
				for k, v := range embeddedMap {
					if _, exists := out[k]; !exists {
						out[k] = v
					}
				}
			}
			continue
		}

		if name == "" {
			name = fieldType.Name
		}

		// 递归处理嵌套结构体和切片
		out[name] = valueToInterface(realVal)
	}

	return out
}

// valueToInterface 将反射值转换为通用接口，如果是结构体或切片则递归处理
func valueToInterface(v reflect.Value) any {
	switch v.Kind() {
	case reflect.Struct:
		return structToMap(v)

	case reflect.Slice, reflect.Array:
		// 递归处理切片中的每个元素
		if v.IsNil() {
			return nil
		}
		length := v.Len()
		sliceOut := make([]any, length)
		for i := 0; i < length; i++ {
			sliceOut[i] = valueToInterface(v.Index(i))
		}
		return sliceOut

	case reflect.Map:
		// 递归处理 Map 的 Value
		if v.IsNil() {
			return nil
		}
		mapOut := make(map[string]any)
		iter := v.MapRange()
		for iter.Next() {
			k := iter.Key()
			mapOut[k.String()] = valueToInterface(iter.Value())
		}
		return mapOut

	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return valueToInterface(v.Elem())

	default:
		// 基础类型直接返回
		return v.Interface()
	}
}
