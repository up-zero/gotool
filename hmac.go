package gotool

import (
	"crypto"
	"crypto/hmac"
)

// hmacGenerate 计算 HMAC
//
// data 数据
// key 密钥
// hash 哈希算法
func hmacGenerate(data, key []byte, hash crypto.Hash) []byte {
	h := hmac.New(hash.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// HmacSHA256 计算 SHA256
//
// data 数据
// key 密钥
func HmacSHA256(data []byte, key []byte) []byte {
	return hmacGenerate(data, key, crypto.SHA256)
}

// HmacSHA384 计算 HmacSHA384
//
// data 数据
// key 密钥
func HmacSHA384(data []byte, key []byte) []byte {
	return hmacGenerate(data, key, crypto.SHA384)
}

// HmacSHA512 计算 HmacSHA512
//
// data 数据
// key 密钥
func HmacSHA512(data []byte, key []byte) []byte {
	return hmacGenerate(data, key, crypto.SHA512)
}
