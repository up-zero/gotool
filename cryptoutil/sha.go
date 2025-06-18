package cryptoutil

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/up-zero/gotool"
	"hash"
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
