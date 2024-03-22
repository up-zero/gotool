package gotool

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
	t.Log(HttpDelete("http://baidu.com", []byte("")))
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
	t.Log(HttpDeleteWithTimeout("http://baidu.com", []byte(""), 10*time.Second))
}
