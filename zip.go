package gotool

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type UnzipNotify struct {
	Progress int    `json:"progress"` // 解压进度
	FPath    string `json:"f_path"`   // 解压文件的路径
}

// Zip 文件夹压缩
//
// dest 压缩到的文件 例如：/var/xxx.zip
// src 源文件夹 例如：/var/xxx
func Zip(dest, src string) error {
	return ZipWithNotify(dest, src, nil)
}

// ZipWithNotify 带通知的文件夹压缩
//
// dest 压缩到的文件 例如：/var/xxx.zip
// src 源文件夹 例如：/var/xxx
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

// Unzip 文件解压
//
// dest 解压到的路径 例如：/var/xxx
// src 文件路径 例如：/var/xxx.zip
func Unzip(dest, src string) error {
	return UnzipWithNotify(dest, src, nil)
}

// UnzipWithNotify 带通知的文件解压
//
// dest 解压到的路径 例如：/var/xxx
// src 文件路径 例如：/var/xxx.zip
// ch 用于通知解压进度
func UnzipWithNotify(dest, src string, ch chan *UnzipNotify) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	index := 0

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path : %s", fpath)
		}
		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// 创建父级文件夹
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		// 打开输出文件
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		// 打开输入文件
		inFile, err := f.Open()
		if err != nil {
			return err
		}

		// 写入文件
		_, err = io.Copy(outFile, inFile)
		outFile.Close()
		inFile.Close()
		if err != nil {
			return err
		}

		// 通知解压进度
		if ch != nil {
			un := &UnzipNotify{
				Progress: index,
				FPath:    fpath,
			}
			ch <- un
			index++
		}
	}
	return nil
}
