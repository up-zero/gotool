package gotool

import "testing"

func TestRunShell(t *testing.T) {
	t.Log(RunShell("ls -al"))
}

func TestRunCommand(t *testing.T) {
	t.Log(RunCommand("ls", "-al"))
	t.Log(RunCommand("go", "version"))
}
