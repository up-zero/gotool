package gotool

import (
	"testing"
)

func TestMd5(t *testing.T) {
	t.Log(Md5("123456"))                           // e10adc3949ba59abbe56e057f20f883e
	t.Log(Md5([]byte("123456")))                   // e10adc3949ba59abbe56e057f20f883e
	t.Log(Md5("e10adc3949ba59abbe56e057f20f883e")) // 14e1b600b1fd579f47433b88e8d85291
}

func TestMd5Iterations(t *testing.T) {
	t.Log(Md5Iterations("123456", 2)) // 14e1b600b1fd579f47433b88e8d85291
}

func TestMd5File(t *testing.T) {
	t.Log(Md5File("LICENSE")) // 79b08c4010a4346486fcc7eb63d7edb1
}
