package gotool

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Zip 文件夹压缩
//
// dest 压缩到的文件 例如：/var/xxx.zip
// src 源文件夹 例如：/var/xxx
func Zip(dest, src string) error {
	return ZipWithNotify(dest, src, nil)
}

// ZipWithNotify 带通知的文件夹压缩
//
// destZip 压缩到的文件 例如：/var/xxx.zip
// srcFile 源文件夹 例如：/var/xxx
// ch 用于通知压缩进度
func ZipWithNotify(dest, src string, ch chan int) error {
	zipFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	index := 0
	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(src)+string(os.PathSeparator))
		if info.IsDir() {
			header.Name += string(os.PathSeparator) // 拼接压缩路径
		} else {
			header.Method = zip.Deflate // 指定压缩算法
		}
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if ch != nil {
				ch <- index
			}
			index++
		}
		return err
	})
	return err
}
