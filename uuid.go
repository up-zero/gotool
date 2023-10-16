package gotool

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"
)

// UUID 生成uuid
func UUID() (string, error) {
	now := time.Now().UnixNano()
	// 生成随机字符串
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// 获取Mac地址
	if Mac != "" {
		interfaces, err := net.Interfaces()
		if err == nil {
			for _, v := range interfaces {
				if len(v.HardwareAddr) >= 6 {
					Mac = v.HardwareAddr.String()
					break
				}
			}
		}
	}
	// 计算时间戳, Mac地址, 随机数的MD5哈希值
	hasher := md5.New()
	_, err = io.WriteString(hasher, fmt.Sprintf("%d%s%s", now, randomBytes, Mac))
	if err != nil {
		return "", err
	}
	hash := hasher.Sum(nil)

	// uuid
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", hash[0:4], hash[4:6], hash[6:8], hash[8:10], hash[10:])
	return uuid, nil
}

// UUIDGenerate UUID 生成
func UUIDGenerate() string {
	uuid, _ := UUID()
	return uuid
}
