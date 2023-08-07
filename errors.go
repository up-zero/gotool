package gotool

import "errors"

// ErrNotSupportType 不支持的类型
var ErrNotSupportType = errors.New("not support type")

// ErrNotIsDir 不是目录
var ErrNotIsDir = errors.New("not is dir")

// ErrDstSrcSame 目标和源是同一个
var ErrDstSrcSame = errors.New("dst and src is same")
