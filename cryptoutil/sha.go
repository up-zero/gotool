package cryptoutil

import (
	"crypto/sha1"
	"fmt"
	"github.com/up-zero/gotool"
)

// Sha1 获取SHA1值
//
// # Params:
//
//	p: 待处理的数据
//	salt: 密码盐
func Sha1(p any, salt ...string) (string, error) {
	var data []byte
	switch v := p.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return "", gotool.ErrNotSupportType
	}
	return sha1Hash(data, salt...), nil
}

func sha1Hash(p []byte, salt ...string) string {
	h := sha1.New()
	h.Write(p)
	for _, s := range salt {
		h.Write([]byte(s))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
