package cryptoutil

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"

	"github.com/up-zero/gotool"
)

type hashFunc func() hash.Hash

// Sha1 获取SHA1值
//
// # Params:
//
//	p: 待处理的数据
//	salt: 密码盐
func Sha1(p any, salt ...string) (string, error) {
	return shaCommon(sha1.New, p, salt...)
}

// Sha256 获取SHA256值
//
// # Params:
//
//	p: 待处理的数据
//	salt: 密码盐
func Sha256(p any, salt ...string) (string, error) {
	return shaCommon(sha256.New, p, salt...)
}

// Sha512 获取SHA512值
//
// # Params:
//
//	p: 待处理数据
//	salt: 密码盐
func Sha512(p any, salt ...string) (string, error) {
	return shaCommon(sha512.New, p, salt...)
}

func shaCommon(fn hashFunc, p any, salt ...string) (string, error) {
	var data []byte
	switch v := p.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return "", gotool.ErrNotSupportType
	}
	var h = fn()
	h.Write(data)
	for _, s := range salt {
		h.Write([]byte(s))
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Sha1File 获取文件SHA1值
func Sha1File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", nil
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
