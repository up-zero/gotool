//go:build windows

package sysutil

import (
	"os/exec"
	"strings"
)

// PsByName 根据程序名查询进程列表
//
// # Params:
//
//	name: 程序名
func PsByName(name string) ([]Process, error) {
	ch := make(chan string)
	defer close(ch)
	finish := make(chan struct{})
	defer close(finish)
	processes := make([]Process, 0)

	go func() {
		for {
			select {
			case line := <-ch:
				// line: "chrome.exe","15712","Console","1","266,760 K"
				fields := strings.Split(line, ",")
				if len(fields) >= 2 && strings.Contains(fields[0], name) {
					ps := Process{
						Pid:  strings.Trim(fields[1], `"`),
						PPid: "",
						Cmd:  strings.Trim(fields[0], `"`),
					}
					processes = append(processes, ps)
				}
			case <-finish:
				return
			}
		}
	}()
	err := ExecShellWithNotify(ch, "tasklist /NH /FO CSV | findstr /I "+name)
	finish <- struct{}{}
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok && exitErr.ExitCode() == 1 {
			return processes, nil
		}
		return nil, err
	}

	return processes, nil
}
