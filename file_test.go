package gotool

import (
	"fmt"
	"testing"
)

func TestFileCopy(t *testing.T) {
	t.Log(FileCopy("LICENSE", "LICENSE_BAK"))
}

func TestFileMove(t *testing.T) {
	t.Log(FileMove("/opt/gotool/test.txt", "/opt/gotool/test/rename.txt"))
}

func TestFileDownload(t *testing.T) {
	t.Log(FileDownload("https://www.baidu.com/img/bd_logo1.png", "baidu.png"))
}

func (dp *DownloadProgress) printProgress() {
	progress := float64(dp.Total) / float64(dp.FileSize) * 100
	fmt.Printf("\rDownloading... %.2f%% complete (%d/%d)", progress, dp.Total, dp.FileSize)
}

func TestFileDownloadWithNotify(t *testing.T) {
	dp := make(chan DownloadProgress)
	go func() {
		for data := range dp {
			data.printProgress()
		}
	}()
	data, err := FileDownloadWithNotify(dp, "https://www.baidu.com/img/bd_logo1.png", "baidu.png")
	if err != nil {
		t.Fatal(err)
	}
	data.printProgress()
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

func TestFileSync(t *testing.T) {
	if err := FileSync("/opt/gotool/test.txt"); err != nil {
		t.Fatal(err)
	}
}
