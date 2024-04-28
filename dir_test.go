package gotool

import "testing"

func TestDirCopy(t *testing.T) {
	t.Log(DirCopy("/Users/getcharzp/repo/github/up-zero/gotool/test/dst", "/Users/getcharzp/repo/github/up-zero/gotool/test/src"))
}

func TestCurrentDirCount(t *testing.T) {
	t.Log(CurrentDirCount("."))
}
