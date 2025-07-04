package netutil

import (
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
