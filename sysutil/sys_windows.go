//go:build windows

package sysutil

func SysUptime() (int64, error) {
	return 0, nil
}
