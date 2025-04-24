//go:build darwin

package sysutil

import (
	"os/exec"
	"regexp"
	"strings"
)

// PsByName 根据程序名查询进程列表
//
// name: 程序名
func PsByName(name string) ([]Process, error) {
	ch := make(chan string)
	defer close(ch)
	finish := make(chan struct{})
	defer close(finish)
	re := regexp.MustCompile(`\s+`)
	processes := make([]Process, 0)

	go func() {
		for {
			select {
			case line := <-ch:
				fields := re.Split(strings.TrimPrefix(line, " "), -1)
				if len(fields) >= 3 {
					ps := Process{
						Pid:  fields[0],
						PPid: "",
						Cmd:  strings.Join(fields[2:], " "),
					}
					processes = append(processes, ps)
				}
			case <-finish:
				return
			}
		}
	}()
	err := ExecShellWithNotify(ch, "pgrep -fl "+name+"")
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
