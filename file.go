package gotool

import (
	"io"
	"net/http"
	"os"
	"path"
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
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
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
