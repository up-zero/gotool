package sysutil

import (
	"bufio"
	"bytes"
	"os/exec"
)

// ExecCommand 运行命令
//
// # Params:
//
//	name: 命令名称
//	arg: 命令参数
func ExecCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return execCmd(cmd, nil)
}

// ExecCommandWithNotify 带通知的运行命令
//
// # Params:
//
//	name: 命令名称
//	arg: 命令参数
func ExecCommandWithNotify(ch chan string, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return execCmd(cmd, ch)
}

// execCmd 执行命令
//
// # Params:
//
//	cmd: 命令
//	ch: 输出通道
func execCmd(cmd *exec.Cmd, ch chan string) error {
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		text := scanner.Text()
		if ch != nil {
			ch <- text
		}
	}
	if err := scanner.Err(); err != nil {
		if ch != nil {
			ch <- err.Error()
		}
		return err
	}

	if err := cmd.Wait(); err != nil {
		if ch != nil {
			ch <- stderr.String()
		}
		return err
	}
	return nil
}
