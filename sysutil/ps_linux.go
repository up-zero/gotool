//go:build linux

package sysutil

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// PsByName 根据程序名查询进程列表
//
// # Params:
//
//	name: 程序名
func PsByName(name string) ([]Process, error) {
	processes := make([]Process, 0)
	name = filepath.Base(name)

	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		pid := file.Name()

		// Read the process command
		cmdPath := filepath.Join("/proc", file.Name(), "cmdline")
		cmdBytes, err := ioutil.ReadFile(cmdPath)
		if err != nil {
			continue
		}

		cmd := string(cmdBytes)
		if strings.Contains(cmd, name) {
			ppid, err := getParentPid(pid)
			if err == nil {
				processes = append(processes, Process{
					Pid:  file.Name(),
					PPid: ppid,
					Cmd:  cmd,
				})
			}
		}
	}
	return processes, nil
}

func getParentPid(pid string) (string, error) {
	statPath := filepath.Join("/proc", pid, "stat")
	statBytes, err := ioutil.ReadFile(statPath)
	if err != nil {
		return "", err
	}

	fields := strings.Fields(string(statBytes))
	if len(fields) < 4 {
		return "", fmt.Errorf("invalid stat file for pid %s", pid)
	}

	return fields[3], nil
}
