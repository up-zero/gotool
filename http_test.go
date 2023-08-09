package gotool

import "testing"

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
