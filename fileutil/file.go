package fileutil

import (
	"encoding/json"
	"github.com/up-zero/gotool"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	DefaultDirPerm  = 0755 // rwxr-xr-x
	DefaultFilePerm = 0644 // rw-r--r--
)

// FileCopy 文件拷贝
//
// # Params:
//
//	src: 源文件路径
//	dst: 目标文件路径
func FileCopy(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 获取源文件权限
	info, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// 确保目标目录存在
	if err := os.MkdirAll(filepath.Dir(dst), DefaultDirPerm); err != nil {
		return err
	}

	// 创建目标文件，使用源文件的权限
	dstFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return nil
}

// FileMove 文件移动
//
// # Params:
//
//	src: 源文件路径
//	dst: 目标文件路径
//
// # Examples:
//
//	FileMove("/opt/gotool/test.txt", "/opt/gotool/test/rename.txt")
func FileMove(src, dst string) error {
	if src == dst {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(dst), DefaultDirPerm); err != nil {
		return err
	}
	err := os.Rename(src, dst)
	if err == nil {
		return nil
	}

	// 移动文件失败时，自动回退到 Copy+Delete
	if err := FileCopy(src, dst); err != nil {
		return err
	}
	return os.Remove(src)
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
// # Params:
//
//	p: 文件路径
//	data: 文件内容
//
// # Examples:
//
//	FileSave("/opt/gotool/test.txt", []byte("hello world"))
//	FileSave("/opt/gotool/test.txt", struct{ Message string }{Message: "hello world"})
func FileSave(p string, data any) error {
	var content []byte
	var err error

	switch v := data.(type) {
	case []byte:
		content = v
	case string:
		content = []byte(v)
	case error:
		content = []byte(v.Error())
	default:
		// 其他类型尝试转 JSON
		content, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
	}

	// 创建目录
	err = os.MkdirAll(filepath.Dir(p), DefaultDirPerm)
	if err != nil {
		return err
	}

	// 写文件
	return os.WriteFile(p, content, DefaultFilePerm)
}

// FileSync 文件同步（将内存中的文件刷新到硬盘中）
//
// # Params:
//
//	filePath: 文件路径
func FileSync(filePath string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	return f.Sync()
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

// Exist 判断文件或目录是否存在
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDir 判断是否是目录
func IsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// IsFile 判断是否是文件
func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
