package gotool

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// FileCopy 文件拷贝
//
// src: 源文件
// dst: 目标文件
func FileCopy(src, dst string) error {
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

// FileMove 文件移动
//
// srcFile 源文件
// dstFile 目标文件
//
// Examples:
//
//	gotool.FileMove("/opt/gotool/test.txt", "/opt/gotool/test/rename.txt")
func FileMove(srcFile, dstFile string) error {
	_, err := os.Stat(srcFile)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dstFile), os.ModePerm); err != nil {
		return err
	}
	if err := os.Rename(srcFile, dstFile); err != nil {
		return err
	}
	return nil
}

type DownloadProgress struct {
	Total    uint64
	FileSize uint64
}

func (dp *DownloadProgress) Write(p []byte) (int, error) {
	n := len(p)
	dp.Total += uint64(n)
	return n, nil
}

// FileDownloadWithNotify 带通知的文件下载
//
// ch 通知进度
// url 文件地址
// filePath 文件路径
func FileDownloadWithNotify(ch chan DownloadProgress, url, filePath string) (*DownloadProgress, error) {
	defer close(ch)
	// 创建文件夹
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return nil, err
	}

	w, err := os.Create(filePath + ".tmp")
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fileSize := resp.ContentLength
	if fileSize <= 0 {
		return nil, fmt.Errorf("invalid Content-Length")
	}

	// 通知进度
	progress := &DownloadProgress{FileSize: uint64(fileSize)}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				select {
				case ch <- *progress:
				default:
				}
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
	defer cancel()

	// 下载文件
	_, err = io.Copy(w, io.TeeReader(resp.Body, progress))
	if err != nil {
		return progress, err
	}
	w.Close()

	// 重命名文件
	if err := os.Rename(filePath+".tmp", filePath); err != nil {
		return progress, err
	}

	return progress, nil
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

// FileSave 保存文件
//
// Examples:
//
//	gotool.FileSave("/opt/gotool/test.txt", []byte("hello world"))
//	gotool.FileSave("/opt/gotool/test.txt", struct{ Message string }{Message: "hello world"})
func FileSave(p string, data any) error {
	var content []byte
	var err error

	// 判断 data 类型
	if reflect.TypeOf(data).Kind() == reflect.Slice && reflect.TypeOf(data).Elem().Kind() == reflect.Uint8 {
		content = data.([]byte)
	} else {
		content, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
	}

	// 创建目录
	err = os.MkdirAll(filepath.Dir(p), os.ModePerm)
	if err != nil {
		return err
	}

	// 写文件
	return os.WriteFile(p, content, os.ModePerm)
}

// FileSync 文件同步（将内存中的文件刷新到硬盘中）
//
// filePath 文件路径
func FileSync(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

// FileRead 读文件（结构体）
//
// filePath 文件路径
// dst 目标结构体
func FileRead(filePath string, dst any) error {
	if reflect.TypeOf(dst).Kind() != reflect.Ptr || reflect.TypeOf(dst).Elem().Kind() != reflect.Struct {
		return ErrDstMustBePointerStruct
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fs.ErrNotExist
	}
	if err := json.Unmarshal(data, dst); err != nil {
		return err
	}
	return nil
}
