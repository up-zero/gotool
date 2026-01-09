package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/up-zero/gotool"
)

// RsaGenerateKey 生成RSA密钥对，返回PEM格式的私钥和公钥字符串
//
// # Params:
//
//	bits: 密钥长度，推荐 2048 或 4096
//
// # Returns:
//
//	prvKey: PEM 格式的私钥字符串
//	pubKey: PEM 格式的公钥字符串
//	err: 错误信息
func RsaGenerateKey(bits int) (prvKey, pubKey string, err error) {
	// private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	derPrivateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	prvBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derPrivateStream,
	}
	prvKey = string(pem.EncodeToMemory(prvBlock))

	// public key
	derPublicStream, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublicStream,
	}
	pubKey = string(pem.EncodeToMemory(pubBlock))

	return prvKey, pubKey, nil
}

// RsaEncrypt RSA 公钥加密
//
// # Params:
//
//	data: 待加密的数据
//	publicKey: PEM 格式的公钥字符串
func RsaEncrypt(data []byte, publicKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, fmt.Errorf("%w, failed parse PEM block", gotool.ErrInvalidParam)
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), data)
}

// RsaDecrypt RSA 私钥解密
//
// # Params:
//
//	ciphertext: 待解密的数据
//	privateKey: PEM 格式的私钥字符串
func RsaDecrypt(ciphertext []byte, privateKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, fmt.Errorf("%w, failed parse PEM block", gotool.ErrInvalidParam)
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
