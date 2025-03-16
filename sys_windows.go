//go:build windows

package gotool

func SysUptime() (int64, error) {
	return 0, nil
}
