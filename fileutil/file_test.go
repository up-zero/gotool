package fileutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestFileCopy(t *testing.T) {
	t.Log(FileCopy("../LICENSE", "../LICENSE_BAK"))
}

func TestFileMove(t *testing.T) {
	t.Log(FileMove("/opt/gotool/test.txt", "/opt/gotool/test/rename.txt"))
}

func TestFileCount(t *testing.T) {
	t.Log(FileCount(".", ".go", ".mod"))
}

func TestFileMainName(t *testing.T) {
	t.Log(FileMainName("/opt/gotool/test.go"))
	t.Log(FileMainName("test.go"))
}

func TestFileSave(t *testing.T) {
	err := FileSave("/opt/gotool/test.txt", "hello world")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileSync(t *testing.T) {
	if err := FileSync("/opt/gotool/test.txt"); err != nil {
		t.Fatal(err)
	}
}

func TestFileRead(t *testing.T) {
	type ss struct {
		Name string `json:"name"`
	}
	// 写文件
	s := ss{
		Name: "test",
	}
	if err := FileSave("/opt/gotool/test.txt", s); err != nil {
		t.Fatal(err)
	}

	// 读文件
	s1 := new(ss)
	if err := FileRead("/opt/gotool/test.txt", s1); err != nil {
		t.Fatal(err)
	} else {
		t.Log(s1)
	}
}

func TestFileSize(t *testing.T) {
	t.Log(FileSize("./file.go"))
}

func TestExist(t *testing.T) {
	testutil.Equal(t, Exist("./file.go"), true)
}

func TestIsDir(t *testing.T) {
	testutil.Equal(t, IsDir("."), true)
}

func TestIsFile(t *testing.T) {
	testutil.Equal(t, IsFile("./file.go"), true)
}
