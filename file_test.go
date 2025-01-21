package gotool

import (
	"fmt"
	"testing"
)

func TestFileCopy(t *testing.T) {
	t.Log(FileCopy("LICENSE_BAK", "LICENSE"))
}

func TestFileMove(t *testing.T) {
	t.Log(FileMove("/opt/gotool/test/rename.txt", "/opt/gotool/test.txt"))
}

func TestFileDownload(t *testing.T) {
	t.Log(FileDownload("https://www.baidu.com/img/bd_logo1.png", "baidu.png"))
}

func TestFileDownloadWithNotify(t *testing.T) {
	dp := make(chan DownloadProgress)
	defer close(dp)
	go func() {
		for data := range dp {
			progress := float64(data.Total) / float64(data.FileSize) * 100
			fmt.Printf("\rDownloading... %.2f%% complete (%d/%d)", progress, data.Total, data.FileSize)
		}
	}()
	t.Log(FileDownloadWithNotify(dp, "https://www.baidu.com/img/bd_logo1.png", "baidu.png"))
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
