package timeutil

import (
	"fmt"
	"time"
)

const (
	DateTime      = "2006-01-02 15:04:05"
	DateTimeMilli = "2006-01-02 15:04:05.000"
	DateOnly      = "2006-01-02"
	TimeOnly      = "15:04:05"

	DateTimeCompact      = "20060102150405"
	DateTimeMilliCompact = "20060102150405000"
	DateCompact          = "20060102"
	TimeCompact          = "150405"

	DateTimeSlash = "2006/01/02 15:04:05"
	DateSlash     = "2006/01/02"

	MonthDay   = "01-02" // 月-日
	HourMinute = "15:04" // 时:分
)

// TransformLayout 时间字符串格式转换
//
// # Params:
//
//	dateTimeStr: 时间字符串
//	srcLayout: 原始时间字符串的格式
//	dstLayout: 目标格式
func TransformLayout(dateTimeStr, srcLayout, dstLayout string) (string, error) {
	if dateTimeStr == "" {
		return "", nil
	}
	t, err := time.ParseInLocation(srcLayout, dateTimeStr, time.Local)
	if err != nil {
		return "", fmt.Errorf("转化日期时间 [%s] 原始字符串格式 [%s] 错误: %w", dateTimeStr, srcLayout, err)
	}
	return t.Format(dstLayout), nil
}

// FormatRFC3339 将 RFC3339 格式转换为指定格式
//
// # Params:
//
//	rfcStr: RFC3339格式的时间字符串
//	targetLayout: 目标格式，默认为 DateTime
func FormatRFC3339(rfcStr string, targetLayout ...string) (string, error) {
	layout := DateTime
	if len(targetLayout) > 0 {
		layout = targetLayout[0]
	}
	t, err := time.Parse(time.RFC3339, rfcStr)
	if err != nil {
		return "", fmt.Errorf("无效的 RFC3339 格式: %w", err)
	}
	return t.Format(layout), nil
}

// FormatRFC1123 将 RFC1123 格式转换为指定格式
//
// # Params:
//
//	rfcStr: RFC1123格式的时间字符串
//	targetLayout: 目标格式，默认为 DateTime
func FormatRFC1123(rfcStr string, targetLayout ...string) (string, error) {
	layout := DateTime
	if len(targetLayout) > 0 {
		layout = targetLayout[0]
	}
	t, err := time.Parse(time.RFC1123, rfcStr)
	if err != nil {
		return "", fmt.Errorf("无效的 RFC1123 格式: %w", err)
	}
	return t.Format(layout), nil
}
