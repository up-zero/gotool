package netutil

import (
	"fmt"
	"testing"
	"time"
)

func TestHttpGet(t *testing.T) {
	t.Log(HttpGet("http://getcharzp.cn"))
}

func TestHttpPost(t *testing.T) {
	t.Log(HttpPost("http://baidu.com", []byte("")))
}

func TestHttpPut(t *testing.T) {
	t.Log(HttpPut("http://baidu.com", []byte("")))
}

func TestHttpDelete(t *testing.T) {
	t.Log(HttpDelete("http://baidu.com"))
}

func TestHttpGetWithTimeout(t *testing.T) {
	t.Log(HttpGetWithTimeout("http://getcharzp.cn", 10*time.Second))
}

func TestHttpPostWithTimeout(t *testing.T) {
	t.Log(HttpPostWithTimeout("http://baidu.com", []byte(""), 10*time.Second))
}

func TestHttpPutWithTimeout(t *testing.T) {
	t.Log(HttpPutWithTimeout("http://baidu.com", []byte(""), 10*time.Second))
}

func TestHttpDeleteWithTimeout(t *testing.T) {
	t.Log(HttpDeleteWithTimeout("http://baidu.com", 10*time.Second))
}

func TestParseResponse(t *testing.T) {
	type resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data"`
	}
	t.Log(ParseResponse[resp](HttpGet("http://192.168.110.253:9000/api/v1/loom/list")))
}

func TestFileDownload(t *testing.T) {
	t.Log(FileDownload("https://www.baidu.com/img/bd_logo1.png", "baidu.png"))
}

func (dp *DownloadProgress) printProgress() {
	progress := float64(dp.Finish) / float64(dp.Total) * 100
	fmt.Printf("\rDownloading... %.2f%% complete (%d/%d)", progress, dp.Finish, dp.Total)
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

func TestFileDownloadWithProgress(t *testing.T) {
	if err := FileDownloadWithProgress("https://www.baidu.com/img/bd_logo1.png", "baidu.png", func(finishBytes, totalBytes int64) {
		fmt.Printf("\rDownloading... %.2f%% complete (%d/%d)", float64(finishBytes)/float64(totalBytes)*100, finishBytes, totalBytes)
	}); err != nil {
		t.Fatal(err)
	}
}
