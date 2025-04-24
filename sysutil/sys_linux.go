//go:build linux

package sysutil

import (
	"github.com/up-zero/gotool"
	"github.com/up-zero/gotool/convertutil"
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
		return 0, gotool.ErrInvalidUptimeFile
	}
	return int64(convertutil.StrToFloat64(fields[0]) * 1000), nil
}
