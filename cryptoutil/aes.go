package cryptoutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// pkcs7Padding 对数据进行 PKCS#7 填充
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pkcs7UnPadding 去除 PKCS#7 填充
func pkcs7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("empty data")
	}
	unpadding := int(origData[length-1])
	if unpadding > length {
		return nil, errors.New("invalid padding")
	}
	return origData[:(length - unpadding)], nil
}

// AesCbcEncrypt 输入明文字符串和密钥，输出 Base64 编码的密文
func AesCbcEncrypt(plainText string, key []byte) (string, error) {
	// 1. 创建 AES 密码器
	// key 的长度必须是 16, 24, 或 32
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()

	// 2. 对明文进行 PKCS#7 填充
	paddedPlaintext := pkcs7Padding([]byte(plainText), blockSize)

	// 3. 创建一个与块大小相等的初始向量 (IV)
	// IV 必须是唯一的，但不必是保密的。通常将其放在密文的前面。
	ciphertext := make([]byte, blockSize+len(paddedPlaintext))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 4. 创建 CBC 模式的加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 5. 执行加密
	mode.CryptBlocks(ciphertext[blockSize:], paddedPlaintext)

	// 6. 将结果进行 Base64 编码
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesCbcDecrypt 输入 Base64 编码的密文和密钥，输出明文字符串
func AesCbcDecrypt(cipherText string, key []byte) (string, error) {
	// 1. Base64 解码
	decodedCiphertext, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// 2. 创建 AES 密码器
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()

	// 3. 检查密文长度是否足够
	if len(decodedCiphertext) < blockSize {
		return "", errors.New("ciphertext too short")
	}

	// 4. 分离 IV 和真正的密文
	iv := decodedCiphertext[:blockSize]
	actualCiphertext := decodedCiphertext[blockSize:]

	// 5. 检查密文长度是否是块大小的整数倍
	if len(actualCiphertext)%blockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	// 6. 创建 CBC 模式的解密器
	mode := cipher.NewCBCDecrypter(block, iv)

	// 7. 执行解密
	decrypted := make([]byte, len(actualCiphertext))
	mode.CryptBlocks(decrypted, actualCiphertext)

	// 8. 去除 PKCS#7 填充
	unpadded, err := pkcs7UnPadding(decrypted)
	if err != nil {
		return "", err
	}

	return string(unpadded), nil
}
