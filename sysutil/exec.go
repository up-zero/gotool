package sysutil

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"sync"
)

// ExecCommand 运行命令
//
// # Params:
//
//	name: 命令名称
//	arg: 命令参数
//
// # Example:
//
//	ExecCommand("go", "version")
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

// ExecCommandWithOutput 执行命令并返回合并后的标准输出和标准错误
//
// # Params:
//
//	name: 命令名称
//	arg: 命令参数
//
// # Example:
//
//	output, err := ExecCommandWithOutput("go", "version")
func ExecCommandWithOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return string(outputBytes), fmt.Errorf("command failed: %w", err)
	}
	return string(outputBytes), nil
}

// ExecCommandWithStream 执行命令并对 stdout 和 stderr 的每一行实时调用回调函数
//
// # Params:
//
//	handler: 回调函数
//	name: 命令名称
//	arg: 命令参数
//
// # Example:
//
//	handler := func(line string, isStderr bool) {
//		fmt.Println("line => ", line, " isStderr => ", isStderr)
//	}
//	if err := ExecCommandWithStream(handler, "go", "version"); err != nil {
//		fmt.Println("exec command error => ", err)
//	}
func ExecCommandWithStream(handler func(line string, isStderr bool), name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	// 分别获取 stdout 和 stderr 的管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			if handler != nil {
				handler(scanner.Text(), false)
			}
		}
	}()
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			if handler != nil {
				handler(scanner.Text(), true)
			}
		}
	}()
	wg.Wait()

	return cmd.Wait()
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
