package sysutil

import (
	"testing"
)

func TestExecShell(t *testing.T) {
	t.Log(ExecShell("ls -al"))
	t.Log(ExecShell("dir"))
}

func TestExecCommand(t *testing.T) {
	t.Log(ExecCommand("ls", "-al"))
	t.Log(ExecCommand("go", "version"))
}

func TestExecCommandWithNotify(t *testing.T) {
	ch := make(chan string)
	finish := make(chan struct{})
	go func() {
		for {
			select {
			case data := <-ch:
				t.Log("data => ", data)
			case <-finish:
				return
			}
		}
	}()
	if err := ExecCommandWithNotify(ch, "go", "version"); err != nil {
		t.Fatal(err)
	}
	finish <- struct{}{}
}

func TestExecShellWithNotify(t *testing.T) {
	ch := make(chan string)
	finish := make(chan struct{})
	go func() {
		for {
			select {
			case data := <-ch:
				t.Log("data => ", data)
			case <-finish:
				return
			}
		}
	}()
	if err := ExecShellWithNotify(ch, "dir"); err != nil {
		t.Fatal(err)
	}
	finish <- struct{}{}
}
