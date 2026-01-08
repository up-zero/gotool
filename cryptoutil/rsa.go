package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/up-zero/gotool"
)

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
