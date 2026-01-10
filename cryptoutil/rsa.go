package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/up-zero/gotool"
	"io"
	"os"
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
	pub, err := parsePublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// RsaDecrypt RSA 私钥解密
//
// # Params:
//
//	ciphertext: 待解密的数据
//	privateKey: PEM 格式的私钥字符串
func RsaDecrypt(ciphertext []byte, privateKey string) ([]byte, error) {
	prv, err := parsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, prv, ciphertext)
}

// RsaEncryptFile 使用 RSA 公钥加密文件
//
// # Params:
//
//	srcPath: 源文件路径
//	dstPath: 加密后的文件输出路径
//	pubKey: PEM 格式公钥字符串
func RsaEncryptFile(srcPath, dstPath, pubKey string) error {
	pub, err := parsePublicKey(pubKey)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 计算每一块的大小
	chunkSize := pub.Size() - 11
	buffer := make([]byte, chunkSize)

	for {
		n, err := srcFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		cipherChunk, err := rsa.EncryptPKCS1v15(rand.Reader, pub, buffer[:n])
		if err != nil {
			return err
		}
		if _, err := dstFile.Write(cipherChunk); err != nil {
			return err
		}
	}

	return nil
}

// RsaDecryptFile 使用 RSA 私钥解密文件
//
// # Params:
//
//	srcPath: 密文文件路径
//	dstPath: 解密后的文件输出路径
//	prvKey: PEM 格式密钥字符串
func RsaDecryptFile(srcPath, dstPath, prvKey string) error {
	prv, err := parsePrivateKey(prvKey)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 解密块的大小
	buffer := make([]byte, prv.Size())

	for {
		n, err := srcFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		plainChunk, err := rsa.DecryptPKCS1v15(rand.Reader, prv, buffer[:n])
		if err != nil {
			return err
		}
		if _, err := dstFile.Write(plainChunk); err != nil {
			return err
		}
	}

	return nil
}

// parsePublicKey 解析公钥
func parsePublicKey(pubKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubKeyStr))
	if block == nil {
		return nil, fmt.Errorf("pem decode error: %w", gotool.ErrInvalidParam)
	}
	// PKIX 解析
	pubAny, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err == nil {
		if pub, ok := pubAny.(*rsa.PublicKey); ok {
			return pub, nil
		}
		return nil, fmt.Errorf("PKIX pub key assert error: %w", gotool.ErrInvalidParam)
	}

	// PKCS1 解析
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err == nil {
		return pub, nil
	}

	return nil, fmt.Errorf("parse PKIX or PKCS1 public key error: %w", gotool.ErrInvalidParam)
}

// parsePrivateKey 解析私钥
func parsePrivateKey(prvKeyStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(prvKeyStr))
	if block == nil {
		return nil, fmt.Errorf("pem decode error: %w", gotool.ErrInvalidParam)
	}

	// PKCS1 解析
	prv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return prv, nil
	}

	// PKCS8 解析
	prvAny, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse PKCS8 private key error: %w", gotool.ErrInvalidParam)
	}
	if prv, ok := prvAny.(*rsa.PrivateKey); ok {
		return prv, nil
	}

	return nil, fmt.Errorf("parse RSA private key error: %w", gotool.ErrInvalidParam)
}
