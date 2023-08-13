package gotool

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
)

// RunShell 运行shell命令或脚本
//
// shell: shell 命令或脚本
func RunShell(shell string) error {
	cmd := exec.Command("/bin/bash", "-c", shell)

	return runCmd(cmd)
}

// RunCommand 运行命令
//
// name: 命令名称
// arg: 命令参数
func RunCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	return runCmd(cmd)
}

func runCmd(cmd *exec.Cmd) error {
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
		log.Println(text)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		log.Println(stderr.String())
		return err
	}
	return nil
}
