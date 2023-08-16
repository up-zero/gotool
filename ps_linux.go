//go:build linux || darwin

package gotool

import (
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
						PPid: fields[1],
						Cmd:  strings.Join(fields[2:], " "),
					}
					processes = append(processes, ps)
				}
			case <-finish:
				return
			}
		}
	}()
	err := ExecShellWithNotify(ch, "ps -e -o pid,ppid,cmd | grep '"+name+"'"+" | grep -v 'grep'")
	finish <- struct{}{}
	if exitErr, ok := err.(*exec.ExitError); ok {
		if exitErr.ExitCode() != 1 {
			return nil, err
		}
	} else {
		return nil, err
	}

	return processes, nil
}
