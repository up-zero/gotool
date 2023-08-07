package gotool

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// Md5 获取md5值
//
// elem: string, []byte
// salt: 加盐
func Md5(elem any, salt ...string) (string, error) {
	switch elem.(type) {
	case string:
		data := elem.(string)
		for _, v := range salt {
			data += v
		}
		return fmt.Sprintf("%x", md5.Sum([]byte(data))), nil
	case []byte:
		data := elem.([]byte)
		for _, v := range salt {
			data = append(data, []byte(v)...)
		}
		return fmt.Sprintf("%x", md5.Sum(data)), nil
	default:
		return "", ErrNotSupportType
	}
}

// Md5Iterations 迭代多次求md5
//
// elem: string, []byte
// iterations: 迭代次数
func Md5Iterations(s any, iterations int) (string, error) {
	var err error
	for i := 0; i < iterations; i++ {
		s, err = Md5(s)
		if err != nil {
			return "", err
		}
	}
	return s.(string), nil
}

// Md5File 获取文件的MD5
//
// path 文件路径
func Md5File(path string) (string, error) {
	src, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer src.Close()

	md5Hash := md5.New()
	_, err = io.Copy(md5Hash, src)
	if err != nil {
		return "", err
	}
	mdByte := md5Hash.Sum(nil)
	return fmt.Sprintf("%x", mdByte), nil
}
