package timeutil

import (
	"time"
)

const (
	DateTime = "2006-01-02 15:04:05"
	DateOnly = "2006-01-02"
	TimeOnly = "15:04:05"
)

// timeFormat 时间格式化
//
// dateTime 时间字符串
// oldLayout 旧的时间格式
// newLayout 新的时间格式
func timeFormat(dateTime, oldLayout, newLayout string) (string, error) {
	t, err := time.ParseInLocation(oldLayout, dateTime, time.Local)
	if err != nil {
		return "", err
	}
	return t.Format(newLayout), nil
}

// RFC3339ToNormalTime RFC3339 日期格式标准化
//
// rfc3339 RFC3339日期格式，如 2006-01-02T15:04:05Z07:00
func RFC3339ToNormalTime(rfc3339 string) (string, error) {
	return timeFormat(rfc3339, time.RFC3339, DateTime)
}

// RFC1123ToNormalTime RFC1123 日期格式标准化
//
// rfc1123 RFC1123日期格式，如 Mon, 02 Jan 2006 15:04:05 MST
func RFC1123ToNormalTime(rfc1123 string) (string, error) {
	return timeFormat(rfc1123, time.RFC1123, DateTime)
}
