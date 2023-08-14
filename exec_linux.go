//go:build linux || darwin

package gotool

// ExecShell 运行shell命令或脚本
//
// shell: shell 命令或脚本
func ExecShell(shell string) error {
	cmd := exec.Command("/bin/bash", "-c", shell)
	return execCmd(cmd, nil)
}

// ExecShellWithNotify 带通知的运行shell命令或脚本
//
// ch: 输出通道
// shell: shell 命令或脚本
func ExecShellWithNotify(ch chan string, shell string) error {
	cmd := exec.Command("/bin/bash", "-c", shell)
	return execCmd(cmd, ch)
}
