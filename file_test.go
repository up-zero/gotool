package gotool

import (
	"testing"
)

func TestFileCopy(t *testing.T) {
	t.Log(FileCopy("LICENSE_BAK", "LICENSE"))
}

func TestFileDownload(t *testing.T) {
	t.Log(FileDownload("https://www.baidu.com/img/bd_logo1.png", "baidu.png"))
}

func TestFileCount(t *testing.T) {
	t.Log(FileCount(".", ".go", ".mod"))
}
