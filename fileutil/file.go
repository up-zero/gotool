package fileutil

import (
	"encoding/json"
	"github.com/up-zero/gotool"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

// FileCopy 文件拷贝
//
// # Params:
//
//	src: 源文件
//	dst: 目标文件
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
// # Params:
//
//	srcFile: 源文件
//	dstFile: 目标文件
//
// # Examples:
//
//	FileMove("/opt/gotool/test.txt", "/opt/gotool/test/rename.txt")
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

// FileCount 获取指定目录下的文件个数
//
// # Params:
//
//	dir: 目录路径
//	suffix: 文件后缀名, 默认为空, 即所有文件
//
// # Examples:
//
//	FileCount("/home/xxx") // 指定目录的文件个数
//	FileCount("/home/xxx", ".jpg") // 指定目录的指定后缀名的文件个数
//	FileCount("/home/xxx", ".jpg", ".png") // 指定目录的多个后缀名的文件个数
func FileCount(dir string, args ...string) (int, error) {
	var cnt = 0
	var suffix = make(map[string]struct{})
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return cnt, err
	}
	if !fileInfo.IsDir() {
		return cnt, gotool.ErrNotIsDir
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
// # Examples:
//
//	FileMainName("/opt/gotool/test.go") // test
//	FileMainName("test.go") // test
func FileMainName(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}

// FileSave 保存文件
//
// # Examples:
//
//	FileSave("/opt/gotool/test.txt", []byte("hello world"))
//	FileSave("/opt/gotool/test.txt", struct{ Message string }{Message: "hello world"})
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
// # Params:
//
//	filePath: 文件路径
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
// # Params:
//
//	filePath: 文件路径
//	dst: 目标结构体
func FileRead(filePath string, dst any) error {
	if reflect.TypeOf(dst).Kind() != reflect.Ptr {
		return gotool.ErrDstMustBePointer
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

// FileSize 获取文件大小
func FileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
