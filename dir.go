package gotool

import (
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
