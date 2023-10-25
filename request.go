package gotool

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
)

// ShouldBindJson json入参绑定
//
// r *http.Request
// data any 待绑定数据结构体指针
func ShouldBindJson(req *http.Request, data any) error {
	if req == nil || req.Body == nil {
		return errors.New("invalid request")
	}
	decode := json.NewDecoder(req.Body)
	if err := decode.Decode(data); err != nil {
		return err
	}
	return nil
}

// ShouldBindQuery query入参绑定
//
// r *http.Request
// data any 待绑定数据结构体指针
func ShouldBindQuery(req *http.Request, data any) error {
	query := req.URL.Query()
	targetType := reflect.TypeOf(data).Elem()
	targetValue := reflect.ValueOf(data).Elem()
	for i := 0; i < targetType.NumField(); i++ {
		fieldType := targetType.Field(i)
		tag := fieldType.Tag.Get("json")
		if tag == "" {
			continue
		}
		fieldValue := targetValue.Field(i)
		param := query.Get(tag)
		if param == "" {
			continue
		}

		switch fieldType.Type.Kind() {
		case reflect.String:
			fieldValue.SetString(param)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			intValue, err := strconv.ParseInt(param, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetInt(intValue)
		case reflect.Float64, reflect.Float32:
			floatValue, err := strconv.ParseFloat(param, 64)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(floatValue)
		default:
			continue
		}
	}
	return nil
}
