package gotool

import (
	"testing"
)

func TestFileCopy(t *testing.T) {
	t.Log(FileCopy("LICENSE_BAK", "LICENSE"))
}

func TestFileMove(t *testing.T) {
	t.Log(FileMove("/opt/gotool/test", "/opt/gotool/test.txt"))
}

func TestFileDownload(t *testing.T) {
	t.Log(FileDownload("https://www.baidu.com/img/bd_logo1.png", "baidu.png"))
}

func TestFileCount(t *testing.T) {
	t.Log(FileCount(".", ".go", ".mod"))
}

func TestFileMainName(t *testing.T) {
	t.Log(FileMainName("/opt/gotool/test.go"))
	t.Log(FileMainName("test.go"))
}

func TestFileSave(t *testing.T) {
	err := FileSave("/opt/gotool/test.txt", []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
}
