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

// AesGcmEncrypt 输入明文字符串和密钥，输出 Base64 编码的密文
func AesGcmEncrypt(plainText string, key []byte) (string, error) {
	// 1. 创建 AES 密码器
	// key 的长度必须是 16, 24, 或 32 字节，以分别对应 AES-128, AES-192, or AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 2. 创建 GCM 实例
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 3. 创建 Nonce
	// Nonce 的长度应该是 GCM 标准的 12 字节。
	// 对于同一个密钥，绝不要重复使用相同的 Nonce。
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 4. 执行加密和认证
	ciphertext := gcm.Seal(nonce, nonce, []byte(plainText), []byte{})

	// 5. 将结果进行 Base64 编码，方便传输
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesGcmDecrypt 输入 Base64 编码的密文和密钥，输出明文字符串。
func AesGcmDecrypt(cipherText string, key []byte) (string, error) {
	// 1. Base64 解码
	decodedCiphertext, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// 2. 创建 AES 密码器和 GCM 实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 3. 分离 Nonce 和真正的密文
	nonceSize := gcm.NonceSize()
	if len(decodedCiphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, actualCiphertext := decodedCiphertext[:nonceSize], decodedCiphertext[nonceSize:]

	// 4. 执行解密和验证
	plaintext, err := gcm.Open(nil, nonce, actualCiphertext, []byte{})
	if err != nil {
		// 常见错误: "cipher: message authentication failed"
		return "", err
	}

	return string(plaintext), nil
}
