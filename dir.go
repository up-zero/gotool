package gotool

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// DirCopy 绝对目录文件拷贝
//
// dst: 目标目录
// src: 源目录
func DirCopy(dst, src string) error {
	// 判断目标和源是否相同
	if strings.TrimSpace(dst) == strings.TrimSpace(src) {
		return ErrDstSrcSame
	}
	// 判断源目录是否存在
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcFileInfo.IsDir() {
		return ErrNotIsDir
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
			err = FileCopy(dstPath, srcPath)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// CurrentDirCount 当前文件夹下(不迭代子文件夹)文件或文件夹的个数
//
// dir 目录路径
//
// Examples:
//
//	gotool.CurrentDirCount("/home/xxx") // 当前文件夹下所有文件及文件夹的个数
//	gotool.CurrentDirCount("/home/xxx", "file") // 当前文件夹下文件的个数
//	gotool.CurrentDirCount("/home/xxx", "dir") // 当前文件夹下文件夹的个数
func CurrentDirCount(dir string, args ...string) (int, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return 0, err
	}
	if !info.IsDir() {
		return 0, ErrNotIsDir
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
