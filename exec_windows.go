//go:build windows

package gotool

import "os/exec"

// ExecShell 运行shell命令或脚本
//
// shell: shell 命令或脚本
func ExecShell(shell string) error {
	cmd := exec.Command("cmd", "/C", shell)
	return execCmd(cmd, nil)
}

// ExecShellWithNotify 带通知的运行shell命令或脚本
//
// ch: 输出通道
// shell: shell 命令或脚本
func ExecShellWithNotify(ch chan string, shell string) error {
	cmd := exec.Command("cmd", "/C", shell)
	return execCmd(cmd, ch)
}
