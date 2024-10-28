package gotool

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// FileCopy 文件拷贝
//
// dst: 目标文件
// src: 源文件
func FileCopy(dst, src string) error {
	// 打开源文件
	reader, err := os.Open(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 创建目标文件夹
	if err = os.MkdirAll(path.Dir(dst), os.ModePerm); err != nil {
		return err
	}
	// 创建目标文件
	writer, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer writer.Close()

	// 拷贝文件
	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}
	return nil
}

// FileDownload 文件下载
//
// url 文件地址
// filePath 文件路径
func FileDownload(url, filePath string) error {
	// 创建文件夹
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	// 创建目录源文件
	writer, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer writer.Close()
	// 下载文件
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 拷贝文件
	if _, err = io.Copy(writer, resp.Body); err != nil {
		return err
	}
	return nil
}

// FileCount 获取指定目录下的文件个数
//
// dir 目录路径
//
// Examples:
//
//	gotool.FileCount("/home/xxx") // 指定目录的文件个数
//	gotool.FileCount("/home/xxx", ".jpg") // 指定目录的指定后缀名的文件个数
//	gotool.FileCount("/home/xxx", ".jpg", ".png") // 指定目录的多个后缀名的文件个数
func FileCount(dir string, args ...string) (int, error) {
	var cnt = 0
	var suffix = make(map[string]struct{})
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return cnt, err
	}
	if !fileInfo.IsDir() {
		return cnt, ErrNotIsDir
	}
	if len(args) > 0 {
		for _, v := range args {
			suffix[v] = struct{}{}
		}
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if len(suffix) == 0 {
				cnt++
			} else {
				if _, ok := suffix[filepath.Ext(path)]; ok {
					cnt++
				}
			}
		}
		return nil
	})
	return cnt, nil
}

// FileMainName 获取指定路径的文件名
//
// filePath 文件路径或文件名
//
// Examples:
//
// gotool.FileMainName("/opt/gotool/test.go") // test
//
// gotool.FileMainName("test.go") // test
func FileMainName(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}
