//go:build linux

package gotool

import (
	"os"
	"strings"
)

// SysUptime 系统启动时间（单位：毫秒）
func SysUptime() (int64, error) {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0, err
	}
	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return 0, ErrInvalidUptimeFile
	}
	return int64(StrToFloat64(fields[0]) * 1000), nil
}
