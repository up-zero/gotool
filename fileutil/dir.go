package fileutil

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/up-zero/gotool"
)

// DirCopy 绝对目录文件拷贝，拷贝 src 文件夹里面的内容到 dst 文件夹中
//
// # Params:
//
//	src: 源目录
//	dst: 目标目录
func DirCopy(src, dst string) error {
	// 判断目标和源是否相同
	if strings.TrimSpace(dst) == strings.TrimSpace(src) {
		return gotool.ErrDstSrcSame
	}
	// 判断源目录是否存在
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcFileInfo.IsDir() {
		return gotool.ErrNotIsDir
	}
	// 创建目标目录
	err = os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	return filepath.Walk(src, func(srcPath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if src == srcPath {
			return nil
		}
		dstPath := strings.Replace(srcPath, src, dst, 1)
		if f.IsDir() {
			err = os.MkdirAll(dstPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			err = FileCopy(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// CurrentDirCount 当前文件夹下(不迭代子文件夹)文件或文件夹的个数
//
// # Params:
//
//	dir: 目录路径
//	fileType: 文件类型, 默认为空, 即所有文件及文件夹, 可选值: file, dir
//
// # Examples:
//
//	CurrentDirCount("/home/xxx") // 当前文件夹下所有文件及文件夹的个数
//	CurrentDirCount("/home/xxx", "file") // 当前文件夹下文件的个数
//	CurrentDirCount("/home/xxx", "dir") // 当前文件夹下文件夹的个数
func CurrentDirCount(dir string, args ...string) (int, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return 0, err
	}
	if !info.IsDir() {
		return 0, gotool.ErrNotIsDir
	}
	var fileType = make(map[string]struct{})
	var cnt = 0
	var batchSize = 100
	if len(args) > 0 {
		for _, v := range args {
			fileType[v] = struct{}{}
		}
	}
	f, err := os.Open(dir)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	for {
		entries, err := f.ReadDir(batchSize)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		if len(fileType) == 0 {
			cnt += len(entries)
			continue
		}
		for _, entry := range entries {
			if _, ok := fileType["file"]; ok && !entry.IsDir() {
				cnt++
			}
			if _, ok := fileType["dir"]; ok && entry.IsDir() {
				cnt++
			}
		}
	}

	return cnt, nil
}

// MkParentDir 创建父级文件夹
//
// filePath 文件路径
func MkParentDir(filePath string) error {
	return os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
}
