package gotool

import "errors"

// ErrNotSupportType 不支持的类型
var ErrNotSupportType = errors.New("not support type")

// ErrNotIsDir 不是目录
var ErrNotIsDir = errors.New("not is dir")

// ErrDstSrcSame 目标和源是同一个
var ErrDstSrcSame = errors.New("dst and src is same")

// ErrInvalidJwtFormat jwt 格式错误
var ErrInvalidJwtFormat = errors.New("invalid jwt format")

// ErrInvalidJwtSignature jwt 签名错误
var ErrInvalidJwtSignature = errors.New("invalid jwt signature")

// ErrExpiredJwt jwt 过期
var ErrExpiredJwt = errors.New("expired jwt")

// ErrInvalidJwtAlgorithm jwt 算法不支持
var ErrInvalidJwtAlgorithm = errors.New("invalid jwt algorithm")

// ErrSrcDstCannotBeNil 源和目标不能为空
var ErrSrcDstCannotBeNil = errors.New("src and dst cannot be nil")

// ErrDstMustBePointerStruct 目标必须是指针结构体
var ErrDstMustBePointerStruct = errors.New("dst must be a pointer struct")

// ErrNotSupportFormat 数据格式不支持
var ErrNotSupportFormat = errors.New("not support format")

// ErrInvalidUptimeFile uptime 文件无效
var ErrInvalidUptimeFile = errors.New("invalid uptime file")
